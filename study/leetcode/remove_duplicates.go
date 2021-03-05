package leetcode

/*
leetcode 习题中
去除连续的重复项
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
*/
func removeDuplicates(nums []int) (class int) {
	// 从0 开始计算
	class = 0

	for x := range nums {
		if x == 0 {
			class++
			continue
		} else {
			if nums[x] == nums[x-1] {
				continue
			} else {
				nums[class] = nums[x]
				class++
			}
		}
	}

	return

}
