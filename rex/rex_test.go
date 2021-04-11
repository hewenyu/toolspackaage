package rex

import (
	"fmt"
	"testing"
	"time"
)

// 程序按照顺序返回
func Test_Start(t *testing.T) {

	input := []int{3, 2, 1}
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go run(i, sleeptime, chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d", endTime.Sub(startTime), len(input))

}

// 超时判断
func Test_After(t *testing.T) {

	input := []int{3, 2, 1}
	timeout := 2
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string)
		go Run(i, sleeptime, timeout, chs[i])
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))

}
