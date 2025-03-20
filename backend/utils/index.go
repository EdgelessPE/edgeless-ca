package utils

import (
	"fmt"
	"math/rand"
)

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// generateRandomCode generates a 6-digit random verification code.
func GenerateRandomCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
