package map_shard_example

import (
	"fmt"
	"github.com/hewenyu/toolspackage/systemctl_plus/random"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys := random.NonceStr()
		
		fmt.Println(keys)
	}
}
