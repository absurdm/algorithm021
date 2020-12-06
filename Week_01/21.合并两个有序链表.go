/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	preHead := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			preHead.Next = l1
			l1 = l1.Next
		} else {
			preHead.Next = l2
			l2 = l2.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	}
	if l2 != nil {
		preHead.Next = l2
	}
	return dummy.Next
}

// @lc code=end

