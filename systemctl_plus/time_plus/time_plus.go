package time_plus

import "time"

const (
	// BaseFormat 标准时间格式处理
	BaseFormat = "2006-01-02 15:04:05"
)

type CTime struct {
	data time.Time
}

func (c *CTime) Get() time.Time {
	return c.data
}
