#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <assert.h>
#include <time.h>

#define MAX_THREADS 12
#define NUM_OPERATIONS 1000000

 
typedef struct __node_t {
    int value;
    struct __node_t *next;
} node_t;

 
typedef struct __queue_t {
    node_t *head;
    node_t *tail;
    pthread_mutex_t head_lock, tail_lock;
} queue_t;

 
void Queue_Init(queue_t *q) {
    node_t *tmp = malloc(sizeof(node_t));
    assert(tmp != NULL);
    tmp->next = NULL;
    q->head = q->tail = tmp;
    pthread_mutex_init(&q->head_lock, NULL);
    pthread_mutex_init(&q->tail_lock, NULL);
}

 
void Queue_Enqueue(queue_t *q, int value) {
    node_t *tmp = malloc (sizeof (node_t));
    assert(tmp != NULL);
    tmp->value = value;
    tmp->next = NULL;

    pthread_mutex_lock(&q->tail_lock);
    q->tail->next = tmp;
    q->tail = tmp;
    pthread_mutex_unlock(&q->tail_lock);
}

 
int  Queue_Dequeue(queue_t *q, int *value) {
    pthread_mutex_lock (&q->head_lock);
    node_t *tmp = q->head;
    node_t *new_head = tmp->next;
    if (new_head == NULL) {
        pthread_mutex_unlock(&q->head_lock);
        return -1; // Queue is empty
    }
    *value = new_head->value;
    q->head = new_head;
    pthread_mutex_unlock(&q->head_lock);
    free(tmp);
    return 0;
}

 
void run_benchmark (int num_threads, void *(*workload)(void *), const char *workload_name) {
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
        int num_threads = thread_counts [i];
        run_benchmark(num_threads,  thread_work, "Mixed Workload");
        run_benchmark(num_threads, enqueue_heavy_work, "Enqueue-Heavy Workload");
        run_benchmark(num_threads, dequeue_heavy_work, "Dequeue-Heavy Workload");
    }
    return 0;
}
