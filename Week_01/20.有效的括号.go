/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	n := len(s)
	// 成对出现才有可能有效
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			// 判断是否匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 删除栈顶元素
			stack = stack[:len(stack)-1]
		}	else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	// stack 为空则匹配成功
	return len(stack) == 0
}
// @lc code=end

