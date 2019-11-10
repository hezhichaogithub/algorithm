// 题目: 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

// 示例:
// 		给定一个链表: 1->2->3->4->5, 和 n = 2.
// 		当删除了倒数第二个节点后，链表变为 1->2->3->5.

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{0, head}
	slow := newHead
	fast := newHead

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return newHead.Next
}