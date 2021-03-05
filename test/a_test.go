package test

import (
	"fmt"
	"testing"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			fmt.Println("x+y")
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestQueue_Start(t *testing.T) {
	c := make(chan int)
	quit := make(chan int, 1)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 10)
			fmt.Println(<-c)
		}

		quit <- 234
	}()
	fibonacci(c, quit)
}

func TestQueue_Start2(t *testing.T) {
	a := make(chan bool, 100)
	b := make(chan bool, 100)
	c := make(chan bool, 100)
	for i := 0; i < 10; i++ {
		a <- true
		b <- true
		c <- true
	}
	for i := 0; i < 10; i++ {
		select {
		case <-a:
			fmt.Println("< a")

		case <-b:
			fmt.Println("< b")

		case <-c:
			fmt.Println("< c")

		default:
			fmt.Println("< default")
		}
	}
}
