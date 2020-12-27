/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (77.31%)
 * Likes:    1044
 * Dislikes: 0
 * Total Accepted:    231.4K
 * Total Submissions: 299.3K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */

// @lc code=start

func permute(nums []int) [][]int {
	result := [][]int{}
	visited := map[int]bool{}

	var dfs func(path []int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			result = append(result, path)
			return
		}
		for _, value := range nums {
			if setStatus, exists := visited[value]; exists && setStatus {
				continue
			}
			visited[value] = true
			dfs(append(path, value))
			visited[value] = false
		}
	}
	dfs([]int{})
	return result
}

// @lc code=end

