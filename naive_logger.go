package main

import (
    "os"
    "log"
)

type NaiveLogger struct {
    file *os.File
}

func NewNaiveLogger(filename string) *NaiveLogger {
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalf("Failed to create file: %v", err)
    }
    return &NaiveLogger{file: file}
}

func (l *NaiveLogger) Log(entry LogEntry) {
    _, err := l.file.WriteString(entry.String() + "\n")
    if err != nil {
        log.Printf("Write error: %v", err)
    }
    err = l.file.Sync()
    if err != nil {
        log.Printf("Sync error: %v", err)
    }
}

func (l *NaiveLogger) Close() {
    l.file.Close()
}
