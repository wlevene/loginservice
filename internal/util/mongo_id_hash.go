package util

import (
	"strconv"
)

/**
 * @brief 对id进行hash，如:member_id, content_id
 */
func MgoHash(key string) uint32 {

	if key == "" {
		return 0
	}

	iKey, err := strconv.Atoi(key)

	if err == nil {
		return (uint32)(iKey)
	}

	return uint32(HashStringToInt(key))

	// return MgoHashV1(key)
}
