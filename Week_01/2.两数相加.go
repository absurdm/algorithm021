/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	headList := &ListNode{}
	head := headList
	num := 0

	for l1 != nil || l2 != nil || num > 0 {
		headList.Next = new(ListNode)
		headList = headList.Next
		if l1 != nil {
			num = num + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num = num + l2.Val
			l2 = l2.Next
		}
		headList.Val = (num) % 10
		num = num / 10
	}

	return head.Next
}

if is bar
// @lc code=end

