// 题目: 反转一个单链表。

// 示例:
// 		输入: 1->2->3->4->5->NULL
// 		输出: 5->4->3->2->1->NULL

// 迭代
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		endNode *ListNode = nil
		cur               = head
	)

	for cur != nil {
		next := cur.Next
		cur.Next = endNode
		endNode = cur
		cur = next
	}

	return endNode
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}