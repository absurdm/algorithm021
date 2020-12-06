/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	return climbStairsDP(n)
}

func climbStairsDP(n int) int {
	if n == 1 {
		return 1
	}
	// 只保存最后结果, 节省空间
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		one, two = two, one+two
	}
	return two
}

// @lc code=end

