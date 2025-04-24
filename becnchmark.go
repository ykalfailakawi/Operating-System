package main

import (
	"fmt"
	"time"
	"math/rand"
)

func benchmarkRAID(label string, raid RAID, totalMB int) {
	blockCount := (totalMB * 1024 * 1024) / blockSize
	 
	block := make([]byte, blockSize)

	 
	rand.Read(block)

	fmt.Printf("Benchmarking %s (%d blocks, %d MB)...\n", label, blockCount, totalMB)

	 
	startWrite := time.Now()
	for i := 0; i < blockCount; i++ {
		err := raid.Write(i, block)
		if err != nil {
			fmt.Printf("Write error at block %d: %v\n", i, err)
			return
		}
	}
	writeDuration := time.Since(startWrite)

	 
	startRead := time.Now()
	for i := 0; i < blockCount; i++ {
		_, err := raid.Read(i)
		if err != nil {
			fmt.Printf("Read error at block %d: %v\n", i, err)
			return
		}
	}
	readDuration := time.Since(startRead)

	fmt.Printf("[%s] Write Time: %v | Read Time: %v\n\n", label, writeDuration, readDuration)
}
