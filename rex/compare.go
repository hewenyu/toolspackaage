package rex

import (
	"fmt"
	"time"
)

const RegEx = "%"

func Run(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", task_id)
		ch <- re
	}
}

func run(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
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
