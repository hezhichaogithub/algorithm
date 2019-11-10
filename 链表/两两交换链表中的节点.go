// 题目: 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

// 示例:
// 		给定 1->2->3->4, 你应该返回 2->1->4->3.

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = swapPairs(next.Next)
	next.Next = head

	return next
}