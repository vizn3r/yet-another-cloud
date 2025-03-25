package util

import "math/rand"

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateHash(n uint) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	return string(b)
}
