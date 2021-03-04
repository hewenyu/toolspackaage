package leetcode

/*
leetcode 习题中
股票最大利润
算法核心：相邻两天，高抛低吸
https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2zsx1/
*/
func maxProfit(prices []int) (max int) {
	// 获利总数值
	max = 0

	if len(prices) == 0 {
		return
	}

	var tmps int

	for x := range prices {
		if (len(prices) - 1) > x {
			tmps = prices[x+1] - prices[x]
			if tmps > 0 {
				max += tmps
			}
		}
	}

	return
}
