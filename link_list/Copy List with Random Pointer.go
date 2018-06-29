package linklist

// https://leetcode.com/explore/learn/card/linked-list/213/conclusion/1229/

/**
 * Definition for singly-linked list with a random pointer.
 * class RandomListNode {
 *     int label;
 *     RandomListNode next, random;
 *     RandomListNode(int x) { this.label = x; }
 * };
 */
func copyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	// 1. 将原节点和拷贝串成list
	// A-B-C-D >> A-a-B-b-C-c-D-d
	iter := head
	var next *RandomListNode
	for iter != nil {
		next = iter.Next
		iter.Next = &RandomListNode{iter.Val, iter.Next, nil}
		iter = next
	}

	// 2.将原来的 random 链接关系复制到拷贝
	// A0-B1-C0-D1 >> A0-a0-B1-b1-C0-c0-D1-d1
	iter = head
	for iter != nil {
		if iter.Random != nil {
			iter.Next.Random = iter.Random.Next
		}
		iter = iter.Next.Next
	}

	// 3.将原list和拷贝分开
	iter = head
	virtualHead := &RandomListNode{}
	copy := virtualHead
	for iter != nil {
		// copy
		copy.Next = iter.Next
		copy = copy.Next
		// original
		iter.Next = iter.Next.Next
		iter = iter.Next
	}

	return virtualHead.Next
}
