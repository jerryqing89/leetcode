package linklist

// https://leetcode.com/explore/learn/card/linked-list/213/conclusion/1228/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// check nil
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// head
	var sum int
	var head, curr *ListNode

	sum = l1.Val + l2.Val
	head = &ListNode{sum % 10, nil}
	curr = head
	sum = sum / 10
	l1 = l1.Next
	l2 = l2.Next

	// func
	getVal := func(n *ListNode) int {
		if n == nil {
			return 0
		}
		return n.Val
	}
	getNext := func(n *ListNode) *ListNode {
		if n == nil {
			return nil
		}
		return n.Next
	}

	// pluse
	for l1 != nil || l2 != nil {
		sum = sum + getVal(l1) + getVal(l2)
		curr.Next = &ListNode{sum % 10, nil}
		sum = sum / 10
		l1 = getNext(l1)
		l2 = getNext(l2)
		curr = curr.Next
	}
	// last place
	if sum > 0 {
		curr.Next = &ListNode{sum, nil}
	}

	return head
}
