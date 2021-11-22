package map_shard_example

import (
	"fmt"
	"github.com/hewenyu/toolspackage/systemctl_plus/random"
	"runtime"
	"testing"
	"time"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys := random.NonceStr()
		
		fmt.Println(keys)
	}
}

type testData struct {
	Data map[int]struct{} `json:"data"`
}

func (t *testData) del(keys int) {
	var newMap = make(map[int]struct{})
	
	for k, v := range t.Data {
		if k != keys {
			newMap[k] = v
		}
		
	}
	t.Data = newMap
	
}

func TestGet(t *testing.T) {
	
	var c = make(chan int)
	
	//
	go func() {
		for {
			c <- 1
		}
		
	}()
	
	for {
		select {
		case <-c:
			printMemStats("now")
		case <-time.After(time.Second * 1):
			fmt.Println(2)
			
		}
	}
}

func printMemStats(mag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%v：分配的内存 = %vKB, GC的次数 = %v\n", mag, m.Alloc/1024, m.NumGC)
}
