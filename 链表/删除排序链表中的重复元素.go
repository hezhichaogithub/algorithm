// 题目: 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

// 示例:
// 		输入: 1->1->2
// 		输出: 1->2

// 		输入: 1->1->2->3->3
// 		输出: 1->2->3

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		cur  = head.Next
		prev = head
	)

	for cur != nil {
		if cur.Val == prev.Val {
			prev.Next = cur.Next
			cur = cur.Next
		} else {
			prev = prev.Next
			cur = cur.Next
		}
	}

	return head
}