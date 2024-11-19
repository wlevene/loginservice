package util

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateSerialNumber() string {
	now := time.Now()
	timestamp := now.Format("20060102150405") // 格式化时间戳，例如：20220328123713
	randomNumber := rand.Intn(10000)          // 生成 0 到 9999 之间的随机数
	serialNumber := timestamp + strconv.Itoa(randomNumber)
	return serialNumber
}
