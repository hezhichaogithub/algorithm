// 题目: 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

// 示例:
// 		给定这个链表：1->2->3->4->5
// 		当 k = 2 时，应当返回: 2->1->4->3->5
// 		当 k = 3 时，应当返回: 3->2->1->4->5

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		i   = 0
		cur = head
	)

	for i < k-1 && cur != nil {
		i = i + 1
		cur = cur.Next
	}

	if cur == nil {
		return head
	}

	next := cur.Next
	cur.Next = nil
	head = reverse(head)
	cur = head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = reverseKGroup(next, k)

	return head
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		cur           = head
		end *ListNode = nil
	)

	for cur != nil {
		next := cur.Next
		cur.Next = end
		end = cur
		cur = next
	}

	return end
}