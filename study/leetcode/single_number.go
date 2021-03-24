package leetcode

import "fmt"

/*
singleNumber只出现一次的数字
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/

任何数和 00 做异或运算，结果仍然是原来的数，即 a⊕0=a。
任何数和其自身做异或运算，结果是 0,即a⊕a=0。
异或运算满足交换律和结合律，即 a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/single-number/solution/zhi-chu-xian-yi-ci-de-shu-zi-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}

	fmt.Print(1 ^ 2)
	return single
}
