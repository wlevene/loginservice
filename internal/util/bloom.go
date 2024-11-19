package util

import (
	"errors"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// https://github.com/zeromicro/go-zero/blob/master/core/bloom/bloom.go
func NewBloomFilter(redis_addr string, key string, bit uint) *bloom.Filter {
	store := redis.New(redis_addr)
	if store == nil {
		return nil
	}

	filter := bloom.New(store, key, bit)
	return filter
}

func BloomAdd(filter *bloom.Filter, key []byte) error {

	if filter == nil {
		return errors.New("bloom filter is nil")
	}
	return filter.Add(key)
}

func BloomExists(filter *bloom.Filter, key []byte) bool {

	if filter == nil {
		return false
	}

	exists, _ := filter.Exists(key)
	return exists
}

func BloomLocal(filter *bloom.Filter, key []byte) bool {

	if filter == nil {
		return false
	}

	exists, _ := filter.Exists(key)
	return exists
}
