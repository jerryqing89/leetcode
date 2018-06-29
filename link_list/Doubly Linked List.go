package linklist

import "fmt"

type MyLinkedList struct {
	Len  int
	Data *DoublyListNode
}

// DoublyListNode for doubly-linked list.
func (l *MyLinkedList) String() (rs string) {
	if l.Data == nil {
		return
	}
	curr := l.Data
	for curr != nil {
		rs = fmt.Sprintf("%s.%d", rs, curr.Val)
		curr = curr.Next
	}
	rs = fmt.Sprintf("[%d]%s", l.Len, rs)
	return
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{0, nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.Len <= index {
		return -1
	}

	cur := this.Data
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Data == nil {
		this.Data = &DoublyListNode{val, nil, nil}
		this.Len++
		return
	}

	cur := &DoublyListNode{
		val,
		nil,
		this.Data,
	}
	this.Data.Prev = cur
	this.Data = cur
	this.Len++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Data == nil {
		this.Data = &DoublyListNode{val, nil, nil}
		this.Len++
		return
	}

	prev := this.Data
	for {
		if prev.Next == nil {
			break
		}
		prev = prev.Next
	}
	prev.Next = &DoublyListNode{val, prev, nil}
	this.Len++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if this.Len < index {
		return
	}
	if this.Len == 0 {
		this.AddAtHead(val)
		return
	}
	prev := this.Data
	for i := 1; i < index; i++ {
		prev = prev.Next
	}
	node := &DoublyListNode{val, prev, prev.Next}
	if prev.Next != nil {
		prev.Next.Prev = node
	}
	prev.Next = node
	this.Len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.Len <= index {
		return
	}
	if this.Len == 1 {
		this.Len--
		this.Data = nil
		return
	}
	if index == 0 {
		this.Data = this.Data.Next
		this.Data.Prev = nil
		this.Len--
		return
	}
	curr := this.Data
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}
	curr.Prev.Next = curr.Next
	this.Len--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
