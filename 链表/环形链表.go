// 题目: 给定一个链表，判断链表中是否有环。

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var (
		fast = head
		slow = head
	)

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}