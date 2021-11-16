package map_struct

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
