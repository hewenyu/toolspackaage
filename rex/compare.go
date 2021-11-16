package rex

import (
	"fmt"
	"time"
)

const RegEx = "%"

func Run(taskId, sleeping, timeout int, ch chan string) {
	chRun := make(chan string)
	go run(taskId, sleeping, chRun)
	select {
	case re := <-chRun:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", taskId)
		ch <- re
	}
}

func run(taskId, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", taskId, sleeptime)
	return
}

// func worker(stopCh <-chan struct{}) {
// 	go func() {
// 		defer fmt.Println("worker exit")
// 		// Using stop channel explicit exit
// 		for {
// 			select {
// 			case <-stopCh:
// 				fmt.Println("Recv stop signal")
// 				return
// 			case <-t.C:
// 			}
// 		}
// 	}()
// 	return
// }
