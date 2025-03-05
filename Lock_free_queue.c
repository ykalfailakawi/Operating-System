#include <stdio.h>
#include <stdlib.h>
#include <stdatomic.h>
#include <pthread.h>
#include <assert.h>
#include <time.h>

#define MAX_THREADS 12
#define NUM_OPERATIONS 1000000

// Node structure  
typedef struct __node_t {
    int value;
    struct __node_t *next;
} node_t;

// Lock-free queue structure  
typedef struct __queue_t {
    _Atomic(node_t *) head;
    _Atomic(node_t *) tail;
} queue_t;

// Initialize queue
void Queue_Init(queue_t *q) {
    node_t *dummy = malloc(sizeof(node_t));
    assert(dummy != NULL);
    dummy->next = NULL;
    atomic_store(&q->head, dummy);
    atomic_store(&q->tail, dummy);
}

 
void Queue_Enqueue(queue_t *q, int value) {
    node_t *new_node = malloc(sizeof(node_t));
    assert(new_node != NULL);
    new_node->value = value;
    new_node->next = NULL;
    
    node_t *tail;
    while (1) {
        tail = atomic_load(&q->tail);
        node_t *next = atomic_load(&tail->next);
        if (tail == atomic_load(&q->tail)) { // Ensure tail to be consistent
            if (next == NULL) { // Queue not in intermediate state
                if (atomic_compare_exchange_weak(&tail->next, &next, new_node)) {
                    atomic_compare_exchange_weak(&q->tail, &tail, new_node);
                    return;
                }
            } else {
                atomic_compare_exchange_weak(&q->tail, &tail, next);
            }
        }
    }
}

 
int Queue_Dequeue (queue_t *q, int *value) {
    node_t *head;
    while (1) {
        head =  atomic_load(&q->head);
        node_t *tail = atomic_load(&q->tail);
        node_t *next = atomic_load(&head->next);
        
        if (head == atomic_load (&q->head)) { // Ensure head is consistent
            if (head == tail) {
                if (next == NULL) {
                    return -1; // Queue is empty
                }
                atomic_compare_exchange_weak(&q->tail, &tail, next);
            } else {
                *value = next->value;
                if (atomic_compare_exchange_weak(&q->head, &head, next)) {
                    free(head);
                    return 0;
                }
            }
        }
    }
}

 
void run_benchmark(int num_threads, void *(*workload)(void *), const char *workload_name) {
    queue_t q;
    Queue_Init(&q);
    pthread_t threads[num_threads];
    struct timespec start, end;
    
    clock_gettime(CLOCK_MONOTONIC, &start);
    for (int i = 0; i < num_threads; i++) {
        pthread_create(&threads[i], NULL, workload, (void *)&q);
    }
    for (int i = 0; i < num_threads; i++) {
        pthread_join(threads[i], NULL);
    }
    clock_gettime(CLOCK_MONOTONIC, &end);
    
    double time_taken = (end.tv_sec - start.tv_sec) + (end.tv_nsec - start.tv_nsec) / 1e9;
    double throughput = (num_threads * (NUM_OPERATIONS / num_threads)) / time_taken;
    printf("%s - Threads: %d, Time: %.6f seconds, Throughput: %.2f ops/sec\n", workload_name, num_threads, time_taken, throughput);
}

void *thread_work(void *arg) {
    queue_t *q = (queue_t *)arg;
    int ops = NUM_OPERATIONS / MAX_THREADS;
    for (int i = 0; i < ops; i++) {
        if (rand() % 2 == 0) {  // Randomized Mixed workload
            Queue_Enqueue(q, i);
        } else {
            int value;
            Queue_Dequeue(q, &value);
        }
    }
    return NULL;
}

void *enqueue_heavy_work(void *arg) {
    queue_t *q = (queue_t *)arg;
    int ops = NUM_OPERATIONS / MAX_THREADS;
    for (int i = 0; i < ops; i++) {
        Queue_Enqueue(q, i);
    }
    return NULL;
}

void *dequeue_heavy_work(void *arg) {
    queue_t *q = (queue_t *)arg;
    int ops = NUM_OPERATIONS / MAX_THREADS;
    for (int i = 0; i < ops; i++) {
        int value;
        Queue_Dequeue(q, &value);
    }
    return NULL;
}

int main() {
    srand(time(NULL));
    int thread_counts[] = {1, 2, 4, 8, 12};
    int num_configs = sizeof(thread_counts) / sizeof(thread_counts[0]);
    
    for (int i = 0; i < num_configs; i++) {
        int num_threads = thread_counts[i];
        run_benchmark(num_threads, thread_work, "Lock-Free Mixed Workload");
        run_benchmark(num_threads, enqueue_heavy_work, "Lock-Free Enqueue-Heavy Workload");
        run_benchmark(num_threads, dequeue_heavy_work, "Lock-Free Dequeue-Heavy Workload");
    }
    return 0;
}
