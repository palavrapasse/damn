package entity

import (
	"crypto/sha256"
	"fmt"
)

type HSHA256 string

func NewHSHA256(plainText string) HSHA256 {
	var h HSHA256

	hashArray := []byte(plainText)

	sum := sha256.Sum256(hashArray)

	encodedString := fmt.Sprintf("%x", sum)

	h = HSHA256(encodedString)

	return h
}
