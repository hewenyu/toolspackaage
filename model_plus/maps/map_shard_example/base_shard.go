package map_shard_example

import "sync"

/* 简单的计数组建,作为分片例子 */

type (
	CountNum struct {
		Keys     string `json:"keys"`      // 相关信息
		Count    int    `json:"count"`     // 总数
		ErrCount int    `json:"err_count"` // 错误数
	}
	
	CountMap struct {
		mu   sync.RWMutex        // 读写锁
		data map[string]CountNum // 这里的string 是md5+时间戳
	}
	
	CountNumList []CountNum // list 数据
)

// NewCountMap 初始化
func NewCountMap() *CountMap {
	
	return &CountMap{
		data: make(map[string]CountNum),
	}
	
}
