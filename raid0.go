package main

import (
	"fmt"
)

 
type RAID interface {
	Write(blockNum int, data []byte) error
	Read(blockNum int) ([]byte, error)
}

 
type RAID0 struct {
	numDisks int
}

/ 
func NewRAID0(numDisks int) *RAID0 {
	return &RAID0{numDisks: numDisks}
}


func (r *RAID0) Write(blockNum int, data []byte) error {
	diskIndex := blockNum % r.numDisks
	diskFile := fmt.Sprintf("disk%d.dat", diskIndex)
	return writeBlock(diskFile, blockNum/r.numDisks, data)
}

 
func (r *RAID0) Read(blockNum int) ([]byte, error) {
	diskIndex := blockNum % r.numDisks
	diskFile := fmt.Sprintf("disk%d.dat", diskIndex)
	return readBlock(diskFile, blockNum/r.numDisks)
}
