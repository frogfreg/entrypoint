package utils

import (
	"crypto/rand"
	"math/big"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes generates a random string with n runes, only letters and numbers
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letterRunes))))
		b[i] = letterRunes[num.Int64()]
	}
	return string(b)
}
