package time_plus

// NewCTime string 转换成 时间
func NewCTime(times string) (CTime, error) {
	
	newTime, err := StrToTime(times)
	
	if err != nil {
		return CTime{}, err
	}
	return CTime{data: newTime}, nil
}

// NewCTimeLocal string 转换成 时间
func NewCTimeLocal(times string) (CTime, error) {
	
	newTime, err := StrToTimeLocal(times)
	
	if err != nil {
		return CTime{}, err
	}
	return CTime{data: newTime}, nil
}
