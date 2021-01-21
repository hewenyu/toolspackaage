package format

/*
stringInSlice 判断是切片中是否存在
*/
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
