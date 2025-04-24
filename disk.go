package main

import (
	"os"
	"fmt"
)

const blockSize = 4096  

 
func writeBlock(filename string, blockNum int, data []byte) error {
	if len(data) > blockSize {
		return fmt.Errorf("data size exceeds block size")
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	offset := int64(blockNum * blockSize)
	_, err = file.Seek(offset, 0)
	if err != nil {
		return err
	}

	 
	padded := make([]byte, blockSize)
	copy(padded, data)

	_, err = file.Write(padded)
	if err != nil {
		return err
	}

	 
	return file.Sync()
}

 
func readBlock(filename string, blockNum int) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	offset :=  int64(blockNum * blockSize)
	_, err =  file.Seek(offset, 0)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, blockSize)
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

