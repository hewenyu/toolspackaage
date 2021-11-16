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
