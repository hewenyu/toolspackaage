package leetcode

/*
rotate 旋转数组 三种答案
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2skh7/
*/
func rotate(nums []int, k int) []int {

	if len(nums) == 0 {
		return nums
	}

	tmps := k % len((nums))

	for x := 0; x < tmps; x++ {
		nums = rotateOne(nums)
	}

	return nums

}

// 只转动一次
func rotateOne(nums []int) []int {

	var tmps int

	tmps = nums[len(nums)-1]

	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}

	nums[0] = tmps

	return nums
}

// 简单直接
func rotateFun1(nums []int, k int) []int {

	if len(nums) == 0 {
		return nums
	}

	var tmpx []int
	var tmps int

	tmpx = make([]int, len(nums))

	tmps = k % len((nums))

	tmpx = append(tmpx, nums[tmps:]...)

	tmpx = append(tmpx, nums[:tmps]...)

	for keys := range tmpx {

		nums[keys] = tmpx[keys]

	}
	return nums
}

/*
rotateFun 这个问题本质上是交换前后区块的问题
因为顺序是没变的,只需要交换前后区块就行了
*/
func rotateFun2(nums []int, k int) []int {

	if len(nums) == 0 {
		return nums
	}

	var tmpx []int
	var tmps int

	tmps = len(nums) - k%len(nums)

	tmpx = make([]int, 0, len(nums))

	tmpx = append(tmpx, nums[tmps:]...)
	tmpx = append(tmpx, nums[:tmps]...)

	for keys := range tmpx {

		nums[keys] = tmpx[keys]

	}
	return nums
}

/*
rotateFun3 第三种解体思路,利用翻转,本质上是前后互换
*/
func rotateFun3(nums []int, k int) []int {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])

	return nums
}

/*
reverse 翻转专用
*/
func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
}
