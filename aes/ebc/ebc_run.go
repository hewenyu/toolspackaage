package ebc

import (
	"crypto/aes"
	"encoding/base64"
)

/*
ECBDecrypter AES ECB 解密
*/
func ECBDecrypter(connext, key string) []byte {
	// base64 解密
	ciphertext, _ := base64.StdEncoding.DecodeString(connext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// ECB mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := NewECBDecrypter(block)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext
}

/*
ECBEncrypter AES ECB 加密
*/
func ECBEncrypter(connext, key string) (msg string) {

	plaintext := []byte(connext)

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, len(plaintext))
	mode := NewECBEncrypter(block)
	mode.CryptBlocks(ciphertext, plaintext)

	msg = base64.StdEncoding.EncodeToString(ciphertext)

	return
}
