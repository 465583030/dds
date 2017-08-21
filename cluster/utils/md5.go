package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5Hash get strings md5 hash
func GetMd5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
