# README - EECE 4811 HW1

## Files
producer_consumer.go -> Producer-Consumer using IPC (pipes) (Upated to include benchmarking for comparison)
goroutine_producer_consumer.go -> Producer-Consumer using Goroutines


## Run Producer-Consumer
1. Compile:
   
   go build -o producer_consumer producer_consumer.go
   
2. Run:
   
   ./producer_consumer -n=1000
   
3. Expected Output:
   
   Execution took: 66.7822ms
  
   

## Run Producer-Consumer
1. Compile and run:
   
   go run goroutine_producer_consumer.go -n=1000
   
 
   
2. Expected Output:
   
   Goroutine Producer-Consumer took: 203.9µs


- GitHub repo link with both `.go` files is submitted.
- I preferred go language as it was recommended and will be helpful in rest of the course
- PART 1 is submitted with the Benchmark results are submitted in Word File
