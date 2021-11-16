package time_plus

// ToString 将时间转换成 string
func (c *CTime) ToString() string {
	return c.data.Format(BaseTimeLineShort)
}
