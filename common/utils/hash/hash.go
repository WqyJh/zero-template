package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 returns the md5 bytes of data.
func Md5(data []byte) []byte {
	digest := md5.New()
	digest.Write(data)
	return digest.Sum(nil)
}

// Md5Hex returns the md5 hex string of data.
func Md5Hex(data []byte) string {
	digest := md5.New()
	digest.Write(data)
	return hex.EncodeToString(digest.Sum(nil))
}
