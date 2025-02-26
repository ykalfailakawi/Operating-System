# README - EECE 4811 HW3

 

Files:
- single_lock.c -> Concurrent linked list using a single global lock.
- Hand-over-Hand.c -> Concurrent linked list using hand-over-hand locking.

Run Basic Version (Single Global Lock):
1. Compile:
   gcc -o single_lock single_lock.c -pthread
2. Run:
   ./single_lock
3. Expected Output:
   Testing Load: 10000 operations per thread with 50% insert ratio
   Threads: 1, Total Time: 0.032825 seconds

Run Hand-Over-Hand Locking:
1. Compile:
   gcc -o Hand-over-Hand Hand-over-Hand.c -pthread
2. Run:
   ./Hand-over-Hand
3. Expected Output:
   Testing Load: 10000 operations per thread with 50% insert ratio
   Threads: 1, Total Time: 0.170421 seconds

 
