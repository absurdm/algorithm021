/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	// 反转三次
	// 1. 所有元素反转
	// 2. 反转前 K 个元素
	// 3. 反转后 k 后面的元素(l-k)个
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}
// @lc code=end

