package aes

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
)

/*
ECBDecrypter AES ECB 解密
*/
func ECBDecrypter(connext, key string) (ciphertext []byte, err error) {
	// base64 解密
	ciphertext, _ = base64.StdEncoding.DecodeString(connext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return
	}

	if len(ciphertext) < aes.BlockSize {

		err = errors.New("ciphertext too short")
		return
	}

	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		err = errors.New("ciphertext is not a multiple of the block size")

		return
	}

	mode := NewECBDecrypter(block)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return
}

/*
ECBEncrypter AES ECB 加密
*/
func ECBEncrypter(connext, key string) (msg string, err error) {

	plaintext := PKCS5Padding([]byte(connext), aes.BlockSize)

	if len(plaintext)%aes.BlockSize != 0 {
		err = errors.New("plaintext is not a multiple of the block size")
		return
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	ciphertext := make([]byte, len(plaintext))
	mode := NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, plaintext)

	msg = base64.StdEncoding.EncodeToString(ciphertext)

	return
}
