/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height)-1
	// 左右两个指针
	// 宽度等于1则左右两个高柱与里面柱子的差的和即为雨水数量
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				// 左面的高柱
				leftMax = height[left]
			} else {
				// 当指针 height[left] < leftMax 且 height[left] < height[right]
				// 当前遍历left的指针的水必然被挡住
				// 水的数量等于 leftMax - 当前指针指向的位置柱子高度(height[left] 的值)
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				// 右面高柱
				rightMax = height[right]
			} else {
				// 当指针 height[right] < rightMax 且 height[right] <= height[left]
				// 等于条件, 最差结果为 0
				// 其余的所有小于的情况必然水被挡住
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}
// @lc code=end

