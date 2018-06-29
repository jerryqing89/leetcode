package linklist

// Singly Linked List
// https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/

type SinglyLinkedList struct {
	Len  int
	Data *ListNode
}

/** Initialize your data structure here. */
func SinglyConstructor() SinglyLinkedList {
	return SinglyLinkedList{0, nil}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *SinglyLinkedList) Get(index int) int {
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
func (this *SinglyLinkedList) AddAtHead(val int) {
	if this.Data == nil {
		this.Data = &ListNode{val, nil}
		this.Len++
		return
	}

	cur := &ListNode{
		val,
		this.Data,
	}
	this.Data = cur
	this.Len++

}

/** Append a node of value val to the last element of the linked list. */
func (this *SinglyLinkedList) AddAtTail(val int) {
	if this.Data == nil {
		this.Data = &ListNode{val, nil}
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
	prev.Next = &ListNode{val, nil}
	this.Len++

}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *SinglyLinkedList) AddAtIndex(index int, val int) {
	if this.Len < index {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	prev := this.Data
	for i := 1; i < index; i++ {
		prev = prev.Next
	}
	cur := &ListNode{
		val,
		prev.Next,
	}
	prev.Next = cur
	this.Len++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *SinglyLinkedList) DeleteAtIndex(index int) {
	if this.Len <= index {
		return
	}

	if index == 0 {
		this.Data = this.Data.Next
		this.Len--
		return
	}

	prev := this.Data
	for i := 1; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	this.Len--
}
