package main

import (
    "fmt"
    "time"
)

type LogEntry struct {
    Timestamp time.Time
    Level     string
    Context   string
    Message   string
}

func (entry LogEntry) String() string {
    return fmt.Sprintf("[%s] [%s] [%s] %s",
        entry.Timestamp.Format("2006-01-02 15:04:05"),
        entry.Level,
        entry.Context,
        entry.Message)
}
