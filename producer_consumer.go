package main

import (
    "fmt"
    "os"
    "os/exec"
    "strconv"
)

func main() {
    // here we created two pipes one is for sending data and other is for synchronization
    dataReader, dataWriter, _ := os.Pipe()  // for number transfer
    syncReader, syncWriter, _ := os.Pipe()  // for sync signal

    // let's fork consumer process
    consumer := exec.Command("/proc/self/exe") // re-executes same program
    consumer.Env = append(os.Environ(), "IS_CONSUMER=1") // consumer identification
    consumer.Stdin = dataReader   // consumer reading
    consumer.Stdout = os.Stdout    // this allow the consumer to print output
    consumer.ExtraFiles = []*os.File{syncWriter} // consumer send ACK via second pipe

    if err := consumer.Start(); err != nil {
        fmt.Println("Error starting consumer process:", err)
        return
    }

    // Producer process sending numbers and waiting for acknowledgment
    for i := 1; i <= 5; i++ {
        fmt.Printf("Producer: %d\n", i)
        dataWriter.Write([]byte(strconv.Itoa(i) + "\n"))

        // here wait for acknowledgment from consumer
        buf := make([]byte, 1)
        syncReader.Read(buf)
    }

    // close the pipes and wait for consumer to finish
    dataWriter.Close()
    consumer.Wait()
}

// this is consumer function
func init() {
    if os.Getenv("IS_CONSUMER") == "1" {
        consumeNumbers()
        os.Exit(0)
    }
}
// this function will consume numbers
func consumeNumbers() {
    var num int
    syncFile := os.NewFile(3, "syncPipe") // Access extra file descriptor for ACK

    for {
        _, err := fmt.Scanf("%d\n", &num)
        if err != nil {
            break // Stop if no data available
        }
        fmt.Printf("Consumer: %d\n", num)

        // Send ack back to producer
        syncFile.Write([]byte{1})
    }
}
