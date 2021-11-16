package time_plus

import "time"

// FormatMinuteLocal 计算当前时间的分钟
func (c *CTime) FormatMinuteLocal() time.Time {
	
	tLocal := c.data.Local()
	mLocal := tLocal.Add(-time.Duration(tLocal.Second() * int(time.Second))) // 减去秒 计算 当前的分钟数
	
	return mLocal
}
