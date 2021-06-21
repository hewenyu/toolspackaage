package server

import (
	"testing"
	"time"
)

func TestQueue_Start(t *testing.T) {
	go ServerDemo()

	for {
		time.Sleep(time.Second)
	}
}
