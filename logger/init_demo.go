package logger

import "fmt"

func init() {
	fmt.Println("初始化")
	SetName("test")
	SetPath("./logs")
	// New()
}
