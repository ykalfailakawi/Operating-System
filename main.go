package main

import (
	"fmt"
)

func main() {
	fmt.Println("RAID Benchmarking Starting...\n")

	 
	raid0 := NewRAID0(5)
	raid1 := NewRAID1(5)
	raid4 := NewRAID4(5)
	raid5 := NewRAID5(5)

	 
	benchmarkRAID("RAID 0", raid0, 100)
	benchmarkRAID("RAID 1", raid1, 100)
	benchmarkRAID("RAID 4", raid4, 100)
	benchmarkRAID("RAID 5", raid5, 100)

	fmt.Println("Benchmarking Complete!")
}
