package util

import (
	"hash/crc32"
)

func UserIDToInt(id string) int {
	return HashStringToInt(id)
}

func HashStringToInt(id string) int {
	bytes := []byte(id)
	hash := crc32.ChecksumIEEE(bytes)
	intHash := int(hash)
	if intHash < 0 {
		intHash = -intHash
	}
	return intHash
}
