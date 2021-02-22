package queue

import (
	"fmt"
	"testing"
	"time"
)

var server = New()

func push() {
	message := "你好"
	for {
		server.Push(message)
		time.Sleep(time.Second * 5)
	}
}

func pop() {
	for {
		resut := server.Pop()
		if resut != nil {
			fmt.Println(resut)
		}
		time.Sleep(time.Second * 1)
	}
}

func TestQueue_Start(t *testing.T) {
	go push()
	go pop()
	for {
		time.Sleep(time.Second)
	}
}
