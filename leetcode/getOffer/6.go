package getOffer

import "go-leetcode/leetcode"

//func reversePrint(head *leetcode.ListNode) []int {
//
//	res := [10000]int{}
//	i := 9999
//	cur := head
//	for cur != nil {
//		res[i] = cur.Val
//		i--
//		cur = cur.Next
//	}
//	return res[i+1:]
//}

/*
递归解法
*/
func reversePrint(head *leetcode.ListNode) []int {
	tmp := []int{}
	reverse(head, &tmp)
	return tmp
}

func reverse(head *leetcode.ListNode, tmp *[]int) {
	if head == nil {
		return
	}
	reverse(head.Next, tmp)
	*tmp = append(*tmp, head.Val)
}
