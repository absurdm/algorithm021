/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 *
 * https://leetcode-cn.com/problems/search-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (39.88%)
 * Likes:    1114
 * Dislikes: 0
 * Total Accepted:    203.3K
 * Total Submissions: 508.7K
 * Testcase Example:  '[4,5,6,7,0,1,2]\n0'
 *
 * 升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为 [4,5,6,7,0,1,2]
 * ）。
 * 
 * 请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [4,5,6,7,0,1,2], target = 0
 * 输出：4
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [4,5,6,7,0,1,2], target = 3
 * 输出：-1
 * 
 * 示例 3：
 * 
 * 
 * 输入：nums = [1], target = 0
 * 输出：-1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10^4 
 * nums 中的每个值都 独一无二
 * nums 肯定会在某个点上旋转
 * -10^4 
 * 
 * 
 */

// @lc code=start
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// left < right 指针位置
	// 指针位置和指针指向的值关系要清晰
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		// 二分查找的分段条件
		// 左侧无反转情况(左侧是升序): 目标值大于等于左侧值 && 目标值小于中间值
		// 左侧有反转(左侧不是升序, 中间有反转点): 目标值小于中间值, 目标值大于最有侧的值, 中间值小于右侧的值
		// 只需要考虑一侧情况, 另外一次不需要考虑
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// OC O(logn)
// SC O(n)

// @lc code=end

