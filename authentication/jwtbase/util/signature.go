package util

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"hash"
)

// Signature 签名
func Signature(data []byte, algorithm string) (string, error) {
	var h hash.Hash = nil
	switch algorithm {
	case "HS256":
		h = sha256.New()
	case "HS384":
		h = sha512.New384()
	case "HS512":
		h = sha512.New()
	default:
		return "", errors.New("unknown algorithm")
	}
	h.Write(data)
	digested := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(digested), nil
}
