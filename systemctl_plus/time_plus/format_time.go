package time_plus

import (
	"errors"
	"strings"
	"time"
)

const (
	BaseTime      = "2006/01/02 15:04:05.000"
	BaseTimeShort = "2006/01/02 15:04:05"
	
	BaseTimeLine      = "2006-01-02 15:04:05.000"
	BaseTimeLineShort = BaseFormat
	
	BaseTimeT  = "2006-01-02T15:04:05"
	BaseTimeTZ = "2006-01-02T15:04:05Z"
)

var TimeSlice = []string{BaseTime, BaseTimeShort, BaseTimeLine, BaseTimeLineShort, BaseTimeT, BaseTimeTZ}

// StrToTimeLocal 支持多种格式
func StrToTimeLocal(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	if len(dateStr) == 0 {
		return time.Time{}, errors.New("数据格式不能为空")
	}
	for _, format := range TimeSlice {
		dateTime, err := time.ParseInLocation(format, strings.TrimSpace(dateStr), time.Local)
		if err != nil {
			continue
		}
		return dateTime, nil
	}
	return time.Time{}, errors.New("没有符合的格式")
}

// StrToTime 支持多种格式
func StrToTime(dateStr string) (time.Time, error) {
	dateStr = strings.TrimSpace(dateStr)
	if len(dateStr) == 0 {
		return time.Time{}, errors.New("数据格式不能为空")
	}
	for _, format := range TimeSlice {
		dateTime, err := time.Parse(format, strings.TrimSpace(dateStr))
		if err != nil {
			continue
		}
		return dateTime, nil
	}
	return time.Time{}, errors.New("没有符合的格式")
}
