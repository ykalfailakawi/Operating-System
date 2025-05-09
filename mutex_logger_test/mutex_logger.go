package main

import (
    "os"
    "sync"
    "log"
)

type MutexLogger struct {
    file      *os.File
    mutex     sync.Mutex
    buffer    []string
    count     int
}

func NewMutexLogger(filename string) *MutexLogger {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalf("Failed to create file: %v", err)
    }
    return &MutexLogger{
        file:   file,
        buffer: make([]string, 0, 10),
    }
}

func (l *MutexLogger) Log(entry LogEntry) {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    l.buffer = append(l.buffer, entry.String())
    l.count++

    if l.count == 10 {
        for _, line := range l.buffer {
            _, err := l.file.WriteString(line + "\n")
            if err != nil {
                log.Printf("Write error: %v", err)
            }
        }
        l.file.Sync()
        l.buffer = l.buffer[:0]
        l.count = 0
    }
}

func (l *MutexLogger) Close() {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    for _, line := range l.buffer {
        l.file.WriteString(line + "\n")
    }
    l.file.Sync()
    l.file.Close()
}
