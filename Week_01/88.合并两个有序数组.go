/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	for m > 0 && n > 0 {
		// 末尾遍历, 将结果插入到 nums1
		// 其中一个数组遍历完成后, 剩余的位置就不需要在调整
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	// 边界条件问题处理 nums 中的元素提前遍历完成
	// nums1 中所有的数据已经都插入到末尾剩余的位置需要将 num2 复制过来
	// nums2 中元素提前结束 nums1 元素不需要变更
	if m == 0 {
		for i:=0;i<n;i++ {
			nums1[i] = nums2[i]
		}
	}
}

// TC O(m+n)
// SC O(1)
// @lc code=end

