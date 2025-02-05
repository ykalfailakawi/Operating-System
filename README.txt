# README - EECE 4811 HW0
## Files
producer_consumer.go -> Producer-Consumer using IPC (pipes)
stack.go -> Stack implementation using array

## Setups to run the file
1. Install Go 
2. Navigate to project folder:
   
   cd ~/my_unix/HW0_OS/part1
   
3. Initialize Go module:
   
   go mod init part1
   

## Run Producer-Consumer
1. Compile:
   
   go build -o producer_consumer producer_consumer.go
   
2. Run:
   
   ./producer_consumer
   
3. Expected Output:
   
   Producer: 1
   Consumer: 1
   Producer: 2
   Consumer: 2
   ...
   Producer: 5
   Consumer: 5
   

## Run Stack Program
1. Execute:
 
   go run stack.go
  
2. Expected Output:
   
   Popped: 30
   Popped: 20
   Popped: 10
   Stack underflow! Nothing to pop
    
## Notes

- GitHub repo link with both `.go` files is submitted.
- I used go language for the two codes files.
