HW7 
Files included:
- main.go          (Main file that runs benchmarks for all RAID levels)
- disk.go          (Handles block-level read/write operations to simulated disks)
- raid0.go         (Implementation of RAID 0)
- raid1.go         (Implementation of RAID 1)
- raid4.go         (Implementation of RAID 4 with dedicated parity)
- raid5.go         (Implementation of RAID 5 with distributed parity)
- benchmark.go     (Benchmarking logic for performance comparison)

Compilation:
No compilation required for Go. Just run all `.go` files using the command below.

Running the program:
```bash
$ go run *.go
