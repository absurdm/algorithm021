/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (53.37%)
 * Likes:    424
 * Dislikes: 0
 * Total Accepted:    158K
 * Total Submissions: 295.7K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出：[2,2]
 * 
 * 
 * 示例 2:
 * 
 * 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出：[4,9]
 * 
 * 
 * 
 * 说明：
 * 
 * 
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
 * 我们可以不考虑输出结果的顺序。
 * 
 * 
 * 进阶：
 * 
 * 
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 * 
 * 
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	// 利用 hashmap 计算重复和判断数据重复
	hash := map[int]int{}
	res :=[]int{}
	for _, num1 := range nums1 {
		if hash[num1] > 0 {
			hash[num1]++
		} else {
			hash[num1] = 1
		}
	}
	for _, num2 := range nums2 {
		if hash[num2] > 0 {
			// 边界条件处理: 利用 nums1 元素构建的 hash table 来进行计算交集和重复数据出现次数
			res = append(res, num2)
			// 处理出现次数的边界条件
			// 判定交集数字次数, 如果 nums2 中次数大于 nums1 中出现次数, 则不再累加一次出现在结果中
			hash[num2]--
		}
	}
	return res
}
// TC = O(m+n)
// SC = O(m+n)
// @lc code=end

