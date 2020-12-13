/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (74.06%)
 * Likes:    121
 * Dislikes: 0
 * Total Accepted:    57.6K
 * Total Submissions: 77.9K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 
 * 
 * 
 * 
 * 
 * 
 * 返回其前序遍历: [1,3,5,6,2,4]。
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// func preorder(root *Node) []int {
// 	if root == nil {
// 		return nil
// 	}
// 	var res []int
// 	res = append(res, root.Val)
// 	// 结束条条件: 当前节点为叶子节点
//     // for 循环为空, 返回
// 	// N叉树的前序遍历, 没有说名左右节点顺序, 从左到右遍历
// 	for _, node := range root.Children {
// 		res = append(res, preorder(node)...)
// 	}
// 	return res
// }

var res []int

func preorder(root *Node) []int {
	res = []int{}
	dfs(root)
	return res
}

func dfs(root *Node) {
	// 切片零值等于 nil, 非 nil 判断
	if root != nil {
		res = append(res, root.Val)
		for _, children := range root.Children {
			dfs(children)
		}
	}
}


// @lc code=end

