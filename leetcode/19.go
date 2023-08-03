package leetcode

/*
*

	删除链表中倒数第N个节点

	dummy 1 2 3 4 5
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{-1, head}
	fast := dummyHead
	slow := dummyHead
	n++
	for n > 0 && fast != nil {
		n--
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
