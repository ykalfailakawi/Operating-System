#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <time.h>

#define MAX_THREADS 8
#define SMALL_LOAD 10000
#define MEDIUM_LOAD 50000
#define LARGE_LOAD 100000
#define INSERT_RATIO 50   

 
typedef struct __node_t {
    int key;
    struct __node_t *next;
    pthread_mutex_t lock;
} node_t;

// Hand-over-Hand Locking Version
typedef struct __list_t {
    node_t *head;
    pthread_mutex_t lock;
} list_t;

 
typedef struct {
    list_t *list;
    int num_operations;
    int insert_ratio;
} thread_args_t;

 
void List_Init(list_t *L) {
    L->head = NULL;
    pthread_mutex_init(&L->lock, NULL);
}

// insert key into the list
void List_Insert(list_t *L, int key) {
    pthread_mutex_lock(&L->lock);
    node_t *new_node = malloc(sizeof(node_t));
    if (new_node == NULL) {
        perror("malloc");
        pthread_mutex_unlock(&L->lock);
        return;
    }
    pthread_mutex_init(&new_node->lock, NULL);
    new_node->key = key;
    new_node->next = L->head;
    L->head = new_node;
    pthread_mutex_unlock(&L->lock);
}

// look up key  
int List_Lookup(list_t *L, int key) {
    pthread_mutex_lock(&L->lock);
    node_t *curr = L->head;
    if (curr) pthread_mutex_lock(&curr->lock);
    pthread_mutex_unlock(&L->lock);
    
    while (curr) {
        if (curr->key == key) {
            pthread_mutex_unlock(&curr->lock);
            return 0; // Found
        }
        node_t *next = curr->next;
        if (next) pthread_mutex_lock(&next->lock);
        pthread_mutex_unlock(&curr->lock);
        curr = next;
    }
    return -1;  
}

 
void List_Free(list_t *L) {
    pthread_mutex_lock(&L->lock);
    node_t *curr = L->head;
    while (curr) {
        node_t *temp = curr;
        curr = curr->next;
        pthread_mutex_destroy(&temp->lock);
        free(temp);
    }
    L->head = NULL;
    pthread_mutex_unlock(&L->lock);
}

void *worker(void *arg) {
    thread_args_t *args = (thread_args_t *)arg;
    list_t *list = args->list;
    int num_operations = args->num_operations;
    int insert_ratio = args->insert_ratio;
    
    struct timespec start, end;
    clock_gettime(CLOCK_MONOTONIC, &start);
    
    for (int i = 0; i < num_operations; i++) {
        if ((rand() % 100) < insert_ratio) {
            List_Insert(list, rand() % 100000);
        } else {
            List_Lookup(list, rand() % 100000);
        }
    }
    
    clock_gettime(CLOCK_MONOTONIC, &end);
    double time_taken = (end.tv_sec - start.tv_sec) + (end.tv_nsec - start.tv_nsec) / 1e9;
    //printf("Thread completed in %.6f seconds\n", time_taken);
    free(arg);
    return NULL;
}

// Benchmark execution time with different loads and insert/lookup ratios
void benchmark() {
    int loads[] = {SMALL_LOAD, MEDIUM_LOAD, LARGE_LOAD};
    int insert_ratios[] = {50, 80, 20}; // Mixed, Insert-heavy, Lookup-heavy
    list_t myList;
    pthread_t threads[MAX_THREADS];
    struct timespec start, end;
    
    for (int l = 0; l < 3; l++) {
        for (int r = 0; r < 3; r++) {
            printf("Testing Load: %d operations per thread with %d%% insert ratio\n", loads[l], insert_ratios[r]);
            List_Init(&myList);
            for (int num_threads = 1; num_threads <= MAX_THREADS; num_threads *= 2) {
                clock_gettime(CLOCK_MONOTONIC, &start);
                for (int i = 0; i < num_threads; i++) {
                    thread_args_t *args = malloc(sizeof(thread_args_t));
                    args->list = &myList;
                    args->num_operations = loads[l];
                    args->insert_ratio = insert_ratios[r];
                    
                    if (pthread_create(&threads[i], NULL, worker, args) != 0) {
                        perror("pthread_create");
                        exit(EXIT_FAILURE);
                    }
                }
                for (int i = 0; i < num_threads; i++) {
                    if (pthread_join(threads[i], NULL) != 0) {
                        perror("pthread_join");
                        exit (EXIT_FAILURE);
                    }
                }
                clock_gettime (CLOCK_MONOTONIC, &end);
                double  time_taken = (end.tv_sec - start.tv_sec) + (end.tv_nsec - start.tv_nsec) / 1e9;
                printf("Threads: %d, Total Time: %.6f seconds\n", num_threads, time_taken);
            }
            List_Free(&myList);
        }
    }
}

int main() {
    srand (time(NULL));
    benchmark();
    return 0;
}
