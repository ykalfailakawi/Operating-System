package main

import (
    "math/rand"
    "sync"
    "time"
    "fmt"
)

var levels = []string{"INFO", "WARN", "ERROR"}

func randomEntry(id int) LogEntry {
    return LogEntry{
        Timestamp: time.Now(),
        Level:     levels[rand.Intn(len(levels))],
        Context:   fmt.Sprintf("req-%03d", id),
        Message:   "Processing request",
    }
}

func main() {
    logger := NewMutexLogger("mutex.log")
    defer logger.Close()

    const numGoroutines = 5
    const entriesPerGoroutine = 50

    start := time.Now()
    var wg sync.WaitGroup
    wg.Add(numGoroutines)

    for i := 0; i < numGoroutines; i++ {
        go func(id int) {
            defer wg.Done()
            for j := 0; j < entriesPerGoroutine; j++ {
                entry := randomEntry(id*100 + j)
                logger.Log(entry)
            }
        }(i)
    }

    wg.Wait()
    elapsed := time.Since(start)
    fmt.Printf("Mutex Logger completed in %s\n", elapsed)
}
