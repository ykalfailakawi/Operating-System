package main

import (
    "flag"
    "fmt"
    "strconv"
    "sync"
    "time"
)

func main() {
    // i have used command line arguments for better benchmarking
    nFlag := flag.String("n", "10000", "Number of messages to exchange")
    flag.Parse()

    // Convert string   to integer
    totalMessages, err := strconv.Atoi(*nFlag)
    if err != nil {
        fmt.Println("Invalid number of messages. Using default (10000).")
        totalMessages = 10000
    }

    startTime := time.Now() // Start timing

    dataChannel := make(chan int) //  
    var wg sync.WaitGroup
    wg.Add(2) // Wait for producer nd consumer

    // Producer goroutine
    go func() {
        defer wg.Done()
        for i := 1; i <= totalMessages; i++ {
            dataChannel <- i
        }
        close(dataChannel) // Close channel when done
    }()

    // consumer goroutine
    go func() {
        defer wg.Done()
        for range dataChannel {  
        }
    }()

    wg.Wait()  
    elapsedTime := time.Since(startTime)
    fmt.Printf("Goroutine Producer-Consumer took: %v\n", elapsedTime)
}
