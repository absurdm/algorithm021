/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	// 快慢指针(双指针)
	// 快指针进行遍历分析旧数组
	// 慢指针用于数据复制
	// 慢指针下标 +1 就是数组的长度(left 值从1开始)
	// TC = O(n)
	// SC = O(1)
	left, right := 1, 1
	for right < n {
		if nums[right] != nums[right-1]{
			// 直接将结果复制到慢指针的下标, 不需要移动元素
			nums[left] = nums[right]
			left++
		}
		right++
	}
	return left
}
// @lc code=end

