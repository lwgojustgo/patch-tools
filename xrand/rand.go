package xrand

import (
	"crypto/md5"
	"hash"
	"math/rand"
)

func Uint64() uint64 {
	return rand.Uint64()
}

func NewMD5() hash.Hash {
	return md5.New()
}

func NewSource(seed int64) rand.Source {
	return rand.NewSource(seed)
}

func Intn(n int) int {
	return rand.Intn(n)
}
