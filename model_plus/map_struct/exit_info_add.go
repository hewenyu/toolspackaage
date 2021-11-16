package map_struct

// Add 新增
func (d *ExitInfo) Add(keys string) {
	
	d.mu.Lock()
	
	defer d.mu.Unlock()
	
	d.data[keys] = struct{}{}
}
