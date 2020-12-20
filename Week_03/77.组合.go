/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 *
 * https://leetcode-cn.com/problems/combinations/description/
 *
 * algorithms
 * Medium (76.09%)
 * Likes:    457
 * Dislikes: 0
 * Total Accepted:    121.3K
 * Total Submissions: 159.4K
 * Testcase Example:  '4\n2'
 *
 * 给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。
 * 
 * 示例:
 * 
 * 输入: n = 4, k = 2
 * 输出:
 * [
 * ⁠ [2,4],
 * ⁠ [3,4],
 * ⁠ [2,3],
 * ⁠ [1,2],
 * ⁠ [1,3],
 * ⁠ [1,4],
 * ]
 * 
 */

// @lc code=start

// func combine(n int, k int) [][]int {
// 	res := [][]int{}
// 	var helper func(start int, path []int)
// 	helper = func(start int, path []int) {
// 		if len(path) == k {
// 			temp := make([]int, len(path))
// 			copy(temp, path)
// 			res = append(res, temp)
// 			return
// 		}

// 		for i := start; i <= n; i++ {
// 			path = append(path, i)
// 			helper(i+1, path)
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	helper(1, []int{})
// 	return res
// }

// 数学概念组合: 从 N 个不同的元素中，任取 K（K ≤ N）个元素为一组
// 还是数学公式强大

// func combine(n int, k int) [][]int {

// 	res := [][]int{}

// 	var helper func(n, k int, path []int)

// 	helper = func(n, k int, path []int) {
// 		if n < k || k == 0 { //
// 			if k == 0 {
// 				temp := make([]int, len(path))
// 				copy(temp, path)
// 				res = append(res, temp)
// 			}
// 			return
// 		}
// 		helper(n-1, k-1, append(path, n))
// 		helper(n-1, k, path)
// 	}
// 	helper(n, k, []int{})
// 	return res
// }

func combine(n int, k int) [][]int {
    ret := [][]int{}
	if n <= 0 || k <= 0 || n < k {
		return ret
	}
	sl := make([]int, k)

	dfs(n, k, 1, sl, &ret)

	return ret
}

func dfs(n, k, start int, sl []int, ret *[][]int) {
    l := len(sl)
	for i := start; i <= n; i++ {
		if k == 1 {
			sl[l-1] = i
			dst := make([]int, l)
			copy(dst, sl)
			*ret = append(*ret, dst)
		} else {
			sl[l-k] = i
			dfs(n, k-1, i+1, sl, ret)
		}
	}
}


// @lc code=end

