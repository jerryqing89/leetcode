package linklist

// https://leetcode.com/explore/learn/card/linked-list/213/conclusion/1227/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// check nil
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	// head
	var head, curr *ListNode
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	curr = head

	// compare
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		} else {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		}
	}

	// merge left
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}

	return head
}
