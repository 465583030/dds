package utils

import (
	"os"
)

// WriteFileBlock write a block to disk at specified point
func WriteFileBlock(path string, block []byte, start int64, end int64) error {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteAt(block, start)
	if err != nil {
		return err
	}

	return nil
}
