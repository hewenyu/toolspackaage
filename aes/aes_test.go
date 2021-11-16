package aes

import (
	"fmt"
	"testing"
)

func Test_ExampleNewECBDecrypter(t *testing.T) {

	key := "11111111111111111111111111111111"
	data := "FbV6Kes1fJ5PtKZwy6IXeoj35RNft9Muz+YUVvwBLP0="

	ciphertext, _ := ECBDecrypt(data, key)

	fmt.Printf("%s\n", ciphertext)

}

func Test_ECBEncrypter(t *testing.T) {
	plaintext := `{"hello":"world"}`

	key := "11111111111111111111111111111111"

	ms, err := ECBEncrypt(plaintext, key)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(ms)

}
