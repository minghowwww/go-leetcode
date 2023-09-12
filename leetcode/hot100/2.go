package hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
每个对应位置数字相加，carry表示进位
结果集每个位置的数字= 两个链表对应位置数字之和 + 进位
如果两个链表长度不一致，短的链表后面的数字默认为0
循环结束后如果仍然有进位，需要在结果集最后添加一个节点
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy1 := &ListNode{}
	dummy1.Next = l1
	dummy2 := &ListNode{}
	dummy2.Next = l2
	p1, p2 := dummy1.Next, dummy2.Next
	res := &ListNode{}
	pr := res

	for p1 != nil || p2 != nil {
		if p1 == nil {
			p1 = &ListNode{Val: 0}
		}
		if p2 == nil {
			p2 = &ListNode{Val: 0}
		}

		sum := p1.Val + p2.Val + carry
		carry = sum / 10

		tmp := &ListNode{Val: sum % 10}
		pr.Next = tmp
		pr = pr.Next

		p1 = p1.Next
		p2 = p2.Next
	}

	if carry > 0 {
		pr.Next = &ListNode{Val: carry}
	}
	return res.Next
}
