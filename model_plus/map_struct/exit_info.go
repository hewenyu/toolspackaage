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

// IsExit 是否存在
func (d *ExitInfo) IsExit(keys string) bool {
	
	d.mu.RLock()
	
	defer d.mu.RUnlock()
	
	_, ok := d.data[keys]
	if ok {
		return true
	} else {
		return false
	}
}

// Add 新增
func (d *ExitInfo) Add(keys string) {
	
	d.mu.Lock()
	
	defer d.mu.Unlock()
	
	d.data[keys] = struct{}{}
}

// Del 删除
func (d *ExitInfo) Del(keys string) {
	
	d.mu.Lock()
	
	defer d.mu.Unlock()
	
	for k, _ := range d.data {
		if k != keys {
			d.data[k] = struct{}{}
		}
	}
}

// ReName 更新
func (d *ExitInfo) ReName(keys string, newKeys string) {
	
	d.mu.Lock()
	
	defer d.mu.Unlock()
	
	for k, _ := range d.data {
		if k == keys {
			d.data[newKeys] = struct{}{}
		}
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
