/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 *
 * https://leetcode-cn.com/problems/climbing-stairs/description/
 *
 * algorithms
 * Easy (51.05%)
 * Likes:    1389
 * Dislikes: 0
 * Total Accepted:    337.2K
 * Total Submissions: 660K
 * Testcase Example:  '2'
 *
 * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
 * 
 * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
 * 
 * 注意：给定 n 是一个正整数。
 * 
 * 示例 1：
 * 
 * 输入： 2
 * 输出： 2
 * 解释： 有两种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶
 * 2.  2 阶
 * 
 * 示例 2：
 * 
 * 输入： 3
 * 输出： 3
 * 解释： 有三种方法可以爬到楼顶。
 * 1.  1 阶 + 1 阶 + 1 阶
 * 2.  1 阶 + 2 阶
 * 3.  2 阶 + 1 阶
 * 
 * 
 */

// @lc code=start
// func climbStairs(n int) int {
// 	if n <= 1 {
// 		return 1
// 	}

// 	return climbStairs(n-1) + climbStairs(n-2)
// }

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

