package main

import (
	"fmt"
)

type RAID1 struct {
	numDisks int
}

 
func NewRAID1(numDisks int) *RAID1 {
	return &RAID1{numDisks: numDisks}
}

 
func (r *RAID1) Write(blockNum int, data []byte) error {
	for i := 0; i < r.numDisks; i++ {
		filename := fmt.Sprintf("disk%d.dat", i)
		err := writeBlock(filename, blockNum, data)
		if err != nil {
			return err
		}
	}
	return nil
}

 
func (r *RAID1) Read(blockNum int) ([]byte, error) {
	filename := "disk0.dat"
	return readBlock(filename, blockNum)
}
