package leetcode

import "github.com/hewenyu/toolspackage/format/slice"

/*
containsDuplicate 判断是否有相同的参数 现排序,后判断是否有相同的
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x248f5/
*/
func ContainsDuplicate(nums []int) bool {

	slice.SortInt(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}

	}

	return false
}

// 利用hash 的方式来处理数据
func ContainsDuplicate_fun2(nums []int) bool {

	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}
