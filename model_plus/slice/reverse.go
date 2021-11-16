package slice

// 尽可能使用原地算法,节约内存使用

// Reverse 对切片进行倒序排列
func Reverse(arr []interface{}) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseString 对切片进行倒序排列
func ReverseString(s []string) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

// ReverseInt 对切片进行倒序排列
func ReverseInt(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseInt8 对切片进行倒序排列
func ReverseInt8(arr []int8) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseInt16 对切片进行倒序排列
func ReverseInt16(arr []int16) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseInt32 对切片进行倒序排列
func ReverseInt32(arr []int32) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseInt64 对切片进行倒序排列
func ReverseInt64(arr []int64) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseUint 对切片进行倒序排列
func ReverseUint(arr []uint) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseUint8 对切片进行倒序排列
func ReverseUint8(arr []uint8) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseUint16 对切片进行倒序排列
func ReverseUint16(arr []uint16) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseUint32 对切片进行倒序排列
func ReverseUint32(arr []uint32) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseUint64 对切片进行倒序排列
func ReverseUint64(arr []uint64) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseUintptr 对切片进行倒序排列
func ReverseUintptr(arr []uintptr) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseFloat32 对切片进行倒序排列
func ReverseFloat32(arr []float32) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseFloat64 对切片进行倒序排列
func ReverseFloat64(arr []float64) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseRune 对切片进行倒序排列
func ReverseRune(s []rune) {
	for from, to := 0, len(s)-1; from < to; from, to = from+1, to-1 {
		s[from], s[to] = s[to], s[from]
	}
}

// ReverseByte 对切片进行倒序排列
func ReverseByte(arr []byte) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseComplex64 对切片进行倒序排列
func ReverseComplex64(arr []complex64) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseComplex128 对切片进行倒序排列
func ReverseComplex128(arr []complex128) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// ReverseBool 对切片进行倒序排列
func ReverseBool(arr []bool) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}

// /*
// reverse 翻转专用,使用原地算法
// */
// func reverse(arr []int) {
// 	for i := 0; i < len(arr)/2; i++ {
// 		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
// 	}
// }
