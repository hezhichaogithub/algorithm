// 题目: 在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

// 示例:
// 		输入: -1->5->3->4->0
// 		输出: -1->0->3->4->5

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		slow = head
		fast = head.Next
	)

	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	mid := slow.Next
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	newHead := &ListNode{0, nil}
	cur := newHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return newHead.Next
}