package linklist

// Classic Problems
// https://leetcode.com/explore/learn/card/linked-list/219/classic-problems/

func reverseList(head *ListNode) *ListNode {
	var prev, next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

func removeElements(head *ListNode, val int) *ListNode {
	first, prev, curr := head, head, head
	for curr != nil {
		// pass over
		if curr.Val != val {
			prev = curr
			curr = curr.Next
			continue
		}

		// head matched
		if first == curr {
			first = curr.Next
		}

		// delete curr
		prev.Next = curr.Next
		curr = curr.Next
	}
	return first
}

func oddEvenList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	odd, even := head, head.Next
	p, q := odd, even
	for q != nil && q.Next != nil {
		p.Next = q.Next
		p = q.Next
		q.Next = p.Next
		q = p.Next
	}
	p.Next = even
	return odd
}

func isPalindrome(head *ListNode) bool {
	// 查找链表中点函数
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// odd:
	if fast != nil {
		// ignore the middle
		slow = slow.Next
	}

	// 反转后半部分
	slow = reverseList(slow)

	// 比较前后值
	fast = head
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}
	return true

}
