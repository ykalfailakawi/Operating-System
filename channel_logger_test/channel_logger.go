package main

import (
    "os"
    "log"
)

type ChannelLogger struct {
    file      *os.File
    logChan   chan LogEntry
    doneChan  chan bool
}

func NewChannelLogger(filename string) *ChannelLogger {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalf("Failed to create file: %v", err)
    }

    logger := &ChannelLogger{
        file:     file,
        logChan:  make(chan LogEntry, 100),
        doneChan: make(chan bool),
    }

    go logger.listen()
    return logger
}

func (l *ChannelLogger) listen() {
    buffer := make([]string, 0, 10)
    for entry := range l.logChan {
        buffer = append(buffer, entry.String())
        if len(buffer) == 10 {
            for _, line := range buffer {
                l.file.WriteString(line + "\n")
            }
            l.file.Sync()
            buffer = buffer[:0]
        }
    }

    // Final flush
    for _, line := range buffer {
        l.file.WriteString(line + "\n")
    }
    l.file.Sync()
    l.doneChan <- true
}

func (l *ChannelLogger) Log(entry LogEntry) {
    l.logChan <- entry
}

func (l *ChannelLogger) Close() {
    close(l.logChan)
    <-l.doneChan
    l.file.Close()
}
