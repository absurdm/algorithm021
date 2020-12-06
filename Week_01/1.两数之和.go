/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	// map 缓存优化, 空间换时间
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			return []int{k, i}
		}
		// 缓存 v 的坐标, 优化一次查找
		m[v] = i
	}
	return nil
}

// @lc code=end

