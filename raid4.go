package main

import (
	"fmt"
)

type RAID4 struct {
	numDisks int  
}

 
func NewRAID4(numDisks int) *RAID4 {
	return &RAID4{numDisks: numDisks}
}

 
func (r *RAID4) Write(blockNum int, data []byte) error {
	if len(data) > blockSize {
		return fmt.Errorf("data too big for block")
	}

	stripeIndex := blockNum / (r.numDisks - 1)
	diskIndex := blockNum % (r.numDisks - 1)
	dataFile := fmt.Sprintf("disk%d.dat", diskIndex)

	// Write actual data block
	err := writeBlock(dataFile, stripeIndex, data)
	if err != nil {
		return err
	}

	 
	parity := make([]byte, blockSize)

	for i := 0; i < r.numDisks-1; i++ {
		file := fmt.Sprintf("disk%d.dat", i)
		block, err := readBlock(file, stripeIndex)
		if err != nil {
			// Treat missing files as zero blocks
			block = make([]byte, blockSize)
		}
		for j := 0; j < blockSize; j++ {
			parity[j] ^= block[j]
		}
	}

	parityDisk := fmt.Sprintf("disk%d.dat", r.numDisks-1)
	return writeBlock(parityDisk, stripeIndex, parity)
}


func (r *RAID4) Read(blockNum int) ([]byte, error) {
	stripeIndex := blockNum / (r.numDisks - 1)
	diskIndex := blockNum % (r.numDisks - 1)
	dataFile := fmt.Sprintf("disk%d.dat", diskIndex)
	return readBlock(dataFile, stripeIndex)
}
