/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (63.83%)
 * Likes:    723
 * Dislikes: 0
 * Total Accepted:    230.2K
 * Total Submissions: 360.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
 * 
 * 
 * 
 * 示例：
 * 二叉树：[3,9,20,null,null,15,7],
 * 
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其层序遍历结果：
 * 
 * 
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 * 
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	bfs(&res, root)
	return res
}

func bfs(res *[][]int, root *TreeNode) {
	if  root ==  nil {
		return
	}

	var queue []int
	queue = append(queue, root.Left.Val)
	queue = append(queue, root.Right.Val)
	*res = append(*res, queue)
	bfs(res, root.Left)
	bfs(res, root.Right)
}

// @lc code=end

