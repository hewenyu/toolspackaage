package time_plus



// NewCTime string 转换成 时间
func NewCTime(times string) (CTime,error) {
	
	tims,err:= StrToTime(times)
	
	if err!= nil{
		return CTime{}, err
	}
	return CTime{data: tims},nil
}

// NewCTimeLocal string 转换成 时间
func NewCTimeLocal(times string) (CTime,error) {
	
	tims,err:= StrToTimeLocal(times)
	
	if err!= nil{
		return CTime{}, err
	}
	return CTime{data: tims},nil
}
