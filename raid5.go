package main

import (
	"fmt"
)

type RAID5 struct {
	numDisks int
}

 
func NewRAID5(numDisks int) *RAID5 {
	return &RAID5{numDisks: numDisks}
}

 
func (r *RAID5) Write(blockNum int, data []byte) error {
	if len(data) > blockSize {
		return fmt.Errorf("data too big for block")
	}

	stripeSize := r.numDisks - 1
	stripeIndex := blockNum / stripeSize
	dataDiskOffset := blockNum % stripeSize

	parityDisk := stripeIndex % r.numDisks
	dataDisk := (dataDiskOffset + (parityDisk + 1)) %  r.numDisks

	 
	err := writeBlock(fmt.Sprintf("disk%d.dat", dataDisk), stripeIndex, data)
	if err != nil {
		return err
	}

	 
	parity := make([]byte, blockSize)
	for i := 0; i < r.numDisks; i++ {
		if i == parityDisk {
			continue
		}
		block, err := readBlock(fmt.Sprintf("disk%d.dat", i), stripeIndex)
		if err != nil {
			block = make([]byte, blockSize)
		}
		for j := 0; j < blockSize; j++ {
			parity[j] ^= block[j]
		}
	}

	 
	return writeBlock(fmt.Sprintf("disk%d.dat", parityDisk), stripeIndex, parity)
}

 

func (r *RAID5) Read(blockNum int) ([]byte, error) {
	stripeSize := r.numDisks -  1
	stripeIndex := blockNum /  stripeSize
	dataDiskOffset := blockNum % stripeSize
	parityDisk := stripeIndex % r.numDisks
	dataDisk := (dataDiskOffset + (parityDisk + 1)) % r.numDisks

	return readBlock( fmt.Sprintf("disk%d.dat", dataDisk), stripeIndex)
}
