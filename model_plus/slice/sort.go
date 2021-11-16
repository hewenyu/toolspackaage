package slice

import "sort"

// SortString 切片排序
func SortString(s []string) {
	sort.Strings(s)
}

// SortInt 切片排序
func SortInt(s []int) {
	sort.Ints(s)
}

// SortInt8 切片排序
func SortInt8(s []int8) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortInt16 切片排序
func SortInt16(s []int16) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortInt32 切片排序
func SortInt32(s []int32) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortInt64 切片排序
func SortInt64(s []int64) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortUint 切片排序
func SortUint(s []uint) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortUint8 切片排序
func SortUint8(s []uint8) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortUint16 切片排序
func SortUint16(s []uint16) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortUint32 切片排序
func SortUint32(s []uint32) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortUint64 切片排序
func SortUint64(s []uint64) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortUintptr 切片排序
func SortUintptr(s []uintptr) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortFloat32 切片排序
func SortFloat32(s []float32) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortFloat64 切片排序
func SortFloat64(s []float64) {
	sort.Float64s(s)
}

// SortRune 切片排序
func SortRune(s []rune) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SortByte 切片排序
func SortByte(s []byte) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}
