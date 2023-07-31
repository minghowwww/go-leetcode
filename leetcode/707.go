package leetcode

/*
*
实现一个单向链表
*/
type MyLinkedList struct {
	head *ListNode
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{&ListNode{-1, nil}, 0}
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}
	pred := this.head
	for index > 0 {
		pred = pred.Next
		index--
	}
	toAdd := &ListNode{val, pred.Next}
	pred.Next = toAdd
	this.size++
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head.Next
	for index > 0 {
		cur = cur.Next
		index--
	}
	return cur.Val
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	cur := this.head
	for index > 0 {
		cur = cur.Next
		index--
	}
	cur.Next = cur.Next.Next
	this.size--
}
