package random

import "math/rand"

// randomString 随机字符串
func randomString(n int) string {
	
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	
	b := make([]rune, n)
	length := len(letter)
	
	for i := range b {
		b[i] = letter[rand.Intn(length)]
	}
	
	return string(b)
}

// NonceStr 随机字符串
func NonceStr() string {
	
	return randomString(32)
}
