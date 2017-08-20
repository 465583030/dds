package utils

import (
	"os"
)

// WriteFileBlock write a block to disk at specified point
func WriteFileBlock(path string, block string, start int64, end int64) error {

	cmd := "touch " + path
	_, err := ExecCmd(cmd)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	file.Seek(start, 0)
	file.Write([]byte(block))

	return nil
}
