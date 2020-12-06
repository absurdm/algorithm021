/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	// 分析情景, 拆分不同情况
	// 只有末尾连续的 9 才会出现进位的情况
	// 进位时当前位置位 0, 高位加一
	// 单全部连续时
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		// 只有9才会导致进位, 其余的均不会出现
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

// @lc code=end

