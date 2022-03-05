package util

import (
	"crypto/rand"
	"math/big"
)

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandString(length int) (string, error) {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := randInt64(0, int64(len(letters)))
		if err != nil {
			return "", err
		}
		result[i] = letters[num]
	}

	return string(result), nil
}

func randInt64(min int64, max int64) (int64, error) {
	numRange := big.NewInt(max - min)
	num, err := rand.Int(rand.Reader, numRange)
	result := num.Int64() + min
	return result, err
}
