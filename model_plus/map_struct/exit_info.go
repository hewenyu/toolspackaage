package map_struct

import "sync"

// ExitInfo 用于判断是否存在信息,一般用于临时的结构体
type ExitInfo struct {
	mu   sync.RWMutex
	data map[string]struct{}
}

// NewExitInfo 初始化
func NewExitInfo() *ExitInfo {
	return &ExitInfo{
		data: make(map[string]struct{}),
	}
}

// GetAll 获取数据
func (d *ExitInfo) GetAll() []string {
	
	d.mu.RLock()
	
	defer d.mu.RUnlock()
	
	keys := make([]string, 0)
	
	for k, _ := range d.data {
		keys = append(keys, k)
	}
	
	return append(keys[:0:0], keys...)
}
