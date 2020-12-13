/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (30.13%)
 * Likes:    2793
 * Dislikes: 0
 * Total Accepted:    373.7K
 * Total Submissions: 1.2M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？请你找出所有满足条件且不重复的三元组。
 * 
 * 注意：答案中不可以包含重复的三元组。
 * 
 * 
 * 
 * 示例：
 * 
 * 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 * 
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	n := len(nums)
	// 使用封装的排序算法
	sort.Ints(nums)

	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		// 去重, 当前坐标与上一个值不同
		if first > 0 && nums[first] == nums[first - 1]{
			continue
		}

		// 从末尾向前的指针
		third := n - 1
		// target 等于自身的负值, a + b + c = 0 == a + b = -c
		target := -1 * nums[first]

		// 避免重复遍历从 first + 1 开始
		for second := first + 1 ; second < n; second ++ {
			// 去重
			if second > first + 1 && nums[second] == nums[second -1] {
				continue
			}

			//开始遍历
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			// 指针指向同一个位置退出
			if second == third {
				break
			}

			if nums[second] + nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}
// @lc code=end

