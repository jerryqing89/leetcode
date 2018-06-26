package linklist

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
