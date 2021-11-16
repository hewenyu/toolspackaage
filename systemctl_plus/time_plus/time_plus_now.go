package time_plus

import "time"

// Now 当前时间
func Now() CTime {
	return CTime{data: time.Now()}
}

// NowLocal 当前时间,使用本地时间
func NowLocal() CTime {
	return CTime{data: time.Now().Local()}
}
