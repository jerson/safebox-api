package util

import (
	"crypto/rand"
	"github.com/rs/xid"
	"math/big"
)

// UniqueID ...
func UniqueID() string {
	guid := xid.New()
	return guid.String()
}

// GenerateRandomASCIIString returns a securely generated random ASCII string.
// It reads random numbers from crypto/rand and searches for printable characters.
// It will return an error if the system's secure random number generator fails to
// function correctly, in which case the caller must not continue.
func GenerateRandomASCIIString(length int) (string, error) {
	result := ""
	for {
		if len(result) >= length {
			return result, nil
		}
		num, err := rand.Int(rand.Reader, big.NewInt(int64(127)))
		if err != nil {
			return "", err
		}
		n := num.Int64()
		// Make sure that the number/byte/letter is inside
		// the range of printable ASCII characters (excluding space and DEL)
		if n > 32 && n < 127 {
			result += string(n)
		}
	}
}