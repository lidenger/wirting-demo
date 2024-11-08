package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
)

var defaultIV = []byte("1234567890123456")

// Encrypt AES/CBC/PKCS#7
func Encrypt(key, data []byte) ([]byte, error) {
	digested := sha256.Sum256(key)
	block, err := aes.NewCipher(digested[:])
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, defaultIV)
	padData := Pad(data, aes.BlockSize)
	cipherText := make([]byte, len(padData))
	mode.CryptBlocks(cipherText, padData)
	return cipherText, nil
}

// Decrypt AES/CBC/PKCS#7
func Decrypt(key, cipherText []byte) ([]byte, error) {
	digested := sha256.Sum256(key)
	block, err := aes.NewCipher(digested[:])
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, defaultIV)
	data := make([]byte, len(cipherText))
	mode.CryptBlocks(data, cipherText)
	originData, err := Unpad(data)
	if err != nil {
		return nil, err
	}
	return originData, nil
}
