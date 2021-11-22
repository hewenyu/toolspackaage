package base_shard

// AddCount 计数功能
func (c CountNum) AddCount(okay bool) CountNum {
	
	if okay {
		return CountNum{
			Keys:     c.Keys,
			Count:    c.Count + 1,
			ErrCount: c.ErrCount,
		}
		
	} else {
		return CountNum{
			Keys:     c.Keys,
			Count:    c.Count + 1,
			ErrCount: c.ErrCount + 1,
		}
	}
}

// IsSet 是否存在
func (c *CountMap) IsSet(keys string) bool {
	
	c.mu.RLock()
	defer c.mu.RUnlock()
	
	_, ok := c.data[keys]
	
	return ok
}

// Set 是否存在
func (c *CountMap) Set(keys string) bool {
	
	c.mu.Lock()
	defer c.mu.Unlock()
	
	_, ok := c.data[keys]
	
	return ok
}
