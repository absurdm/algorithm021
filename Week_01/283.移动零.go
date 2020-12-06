/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// i 快指针遍历
	// j 慢指针将非零值记录复制到慢指针位置
	// 快指针最慢遍历的速度也会和 j 一样
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			// 遇到非零元素慢指针位置加 1
			j++
		}
	}
}
// TC = O(n)
// SC = O(1)
// @lc code=end

