package linklist

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// DoublyListNode for doubly-linked list.
type DoublyListNode struct {
	Val  int
	Next *DoublyListNode
	Prev *DoublyListNode
}

// RandomListNode for random linked list.
type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}
