package linklist

// Two Pointer Technique
// https://leetcode.com/explore/learn/card/linked-list/214/two-pointer-technique/

// HasCycle Linked List Cycle
func (this *ListNode) HasCycle() bool {
	slow := this
	fast := this

	for {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
}

// DetectCycle   Linked List Cycle II
func (this *ListNode) DetectCycle() *ListNode {

	slow := this
	fast := this

	// 1. whether has cycle
	hasCycle := false

	for {
		if fast == nil || fast.Next == nil {
			break
		}

		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			hasCycle = true
			break
		}
	}

	if hasCycle == false {
		return nil
	}

	// 2. find out start point

	// x 非环长  y 弧长
	// n 某整数  C 圆周
	// t 第一次相遇时间

	//   2x+2y=2t
	//   x+y+nC=2t
	// so： x=nC-y
	// slow move to head, repeate x both slow and fast
	//   x+y+nC+(nC-y)=x+2nC
	// both slow and fast at start point

	slow = this
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// GetIntersectionNode Intersection of Two Linked Lists
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lA, lB := 0, 0
	pA, pB := headA, headB

	// 各自遍历得到 tail 和 length
	for pA.Next != nil {
		pA = pA.Next
		lA++
	}
	for pB.Next != nil {
		pB = pB.Next
		lB++
	}
	// 判断共同 tail
	if pA != pB {
		return nil
	}

	// 削平更长链表多的节点
	pA = headA
	pB = headB
	if lA >= lB {
		for lA > lB {
			pA = pA.Next
			lA--
		}
	} else {
		for lA < lB {
			pB = pB.Next
			lB--
		}
	}
	// 寻找交叉点
	for pA != pB {
		pA = pA.Next
		pB = pB.Next
	}
	return pA
}

// RemoveNthFromEnd Intersection of Two Linked Lists
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	// new start
	start := &ListNode{0, nil}
	start.Next = head

	// gen slow and fast with n gap
	slow, fast := start, start
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// move forward
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// delete Nth from end
	slow.Next = slow.Next.Next

	return start.Next
}
