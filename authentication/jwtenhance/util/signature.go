package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
)

// Signature 签名
func Signature(data, key []byte, algorithm string) ([]byte, error) {
	var h func() hash.Hash = nil
	switch algorithm {
	case "HS256":
		h = sha256.New
	case "HS384":
		h = sha512.New384
	case "HS512":
		h = sha512.New
	default:
		return nil, errors.New("unknown algorithm")
	}
	hm := hmac.New(h, key)
	hm.Write(data)
	digested := hm.Sum(nil)
	return digested, nil
}
