package leetcode

/*
*
判断链表是否有环，如果有环返回环的入口 pos

快慢指针 解法
*/
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head
	pos := &ListNode{-1, nil}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			// 有环
			mPoint := fast
			sPoint := head
			for sPoint != mPoint {
				sPoint = sPoint.Next
				mPoint = mPoint.Next
			}
			pos = sPoint
			return pos
		}
	}
	return nil
}
