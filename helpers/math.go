package helpers

import (
	"math/rand"
	"time"
)

// RandRange 获取指定范围内的可变随机整数数，正负都行。[a, b]
func RandRange(_min int64, _max int64) int64 {
	var _rand int64
	if _min >= _max {
		_rand = 0
	} else {
		rand.Seed(time.Now().UnixNano())
		_rand = rand.Int63n(_max-_min) + _min
	}
	return _rand
}

// RandString 生成指定长度的字符串
func RandString(_length int64) string {
	var length int64
	if _length >= 1 {
		length = _length
	} else {
		length = 1
	}
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < int(length); i++ {
		result = append(result, bytes[r.Int63n(int64(len(bytes)))])
	}
	return string(result)
}
