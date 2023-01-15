package entity

import (
	"encoding/base64"
)

type Base64 string

func NewBase64(plainText string) Base64 {
	return Base64(base64.StdEncoding.EncodeToString([]byte(plainText)))
}
