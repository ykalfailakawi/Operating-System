# HW8: Concurrent Logging in Go

This project implements and compares three concurrent loggers in Go:

Naive Logger: No synchronization, fsync after every write
Mutex Logger: Uses sync.Mutex with batched fsync
Channel Logger: Uses chan LogEntry and a dedicated writer goroutine

Directory Structure

naive_logger: Naive logger source code
mutex_logger_test: Mutex logger source code
channel_logger_test: Channel logger source code
 
## How to Run

Each subfolder contains a main.go. Run any logger using:

 
cd naive_logger   
go run .