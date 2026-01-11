package common

import (
	"crypto/rand"
	"math/big"
)

func Random10Digit() int64 {
	min := int64(1_000_000_000) // 10 หลัก
	max := int64(9_999_999_999)

	n, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
	return n.Int64() + min
}
