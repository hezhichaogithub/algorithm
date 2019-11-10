// 题目: 编写一个程序，找到两个单链表相交的起始节点。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var (
		pa = headA
		pb = headB
	)

	for {
		for pa != nil && pb != nil {
			if pa == pb {
				return pa
			}
			pa = pa.Next
			pb = pb.Next
		}

		if pa == nil && pb == nil {
			return nil
		}

		if pa == nil {
			pa = headB
		}
		if pb == nil {
			pb = headA
		}
	}
}