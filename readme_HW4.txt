# README - EECE 4811/5811 HW4

Files:

two_lock_queue.c -> Two-Lock Queue (head_lock & tail_lock).

lock_free_queue.c -> Lock-Free Queue (atomic CAS).

Run Benchmarks
Two-Lock Queue:
gcc -o two_lock_queue two_lock_queue.c -pthread
./two_lock_queue

Lock-Free Queue:
gcc -o lock_free_queue lock_free_queue.c -pthread
./lock_free_queue

Expected Output (Sample):
Mixed Workload - Threads: 1, Time: 0.003206s, Throughput: 311934618 ops/sec
Lock-Free Mixed Workload - Threads: 1, Time: 0.003250s, Throughput: 307711149 ops/sec
...

Benchmark Details:

Workloads: Mixed (50% Enqueue/Dequeue), Enqueue-Heavy, Dequeue-Heavy.

Thread Counts: 1, 2, 4, 8, 12

Metrics: Execution time & Throughput