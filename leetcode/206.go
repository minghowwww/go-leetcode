package leetcode

/*
*
翻转链表
双指针
*/
//func reverseList(head *ListNode) *ListNode {
//	var pre *ListNode
//	cur := head
//	for cur != nil {
//		tmp := cur.Next
//		cur.Next = pre
//
//		pre = cur
//		cur = tmp
//	}
//	return pre
//}

/*
*
递归
*/
func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre *ListNode, cur *ListNode) *ListNode {
	if cur == nil {
		return pre
	}
	tmp := cur.Next
	cur.Next = pre
	return reverse(cur, tmp)
}
