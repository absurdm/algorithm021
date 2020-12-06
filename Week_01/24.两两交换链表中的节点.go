f/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 根据 判断指针交换的条件
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	for head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		prev.Next = next
		prev = head
		head = head.Next
	}
	return dummy.Next
}

// @lc code=end

