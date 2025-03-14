package utils

import (
	"encoding/hex"

	"github.com/zeebo/blake3"
)

func HashStringToHexBlake3(input string) string {
	hasher := blake3.New()
	hasher.Write([]byte(input))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}
