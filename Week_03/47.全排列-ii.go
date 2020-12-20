/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (62.35%)
 * Likes:    545
 * Dislikes: 0
 * Total Accepted:    126.6K
 * Total Submissions: 203.1K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10 
 * 
 * 
 */

// @lc code=start
// 抄的题解
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	result := [][]int{}
	sort.Ints(nums)
	helper(nums, 0, &result)

	return result
}

// 回溯函数实现
// i表示本次函数需要放置的元素位置
func helper(nums []int, i int, result *[][]int) {
	n := len(nums)
	if i == n-1 {
		tmp := make([]int, n)
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字
		if k != i && nums[k] == nums[i] {
			continue
		}
		nums[k], nums[i] = nums[i], nums[k]
		helper(nums, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		nums[i], nums[k] = nums[k], nums[i]
	}
}

// @lc code=end

