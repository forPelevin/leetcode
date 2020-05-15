package linkedlist

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/603/
//Remove Nth Node From End of List
//Solution
//Given a linked list, remove the n-th node from the end of list and return its head.
//
//Example:
//
//Given linked list: 1->2->3->4->5, and n = 2.
//
//After removing the second node from the end, the linked list becomes 1->2->3->5.
//Note:
//
//Given n will always be valid.
//
//Follow up:
//
//Could you do this in one pass?

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n < 1 {
		return head
	}

	if n == 1 {
		if head.Next == nil {
			return nil
		}

		i := head
		for i.Next.Next != nil {
			i = i.Next
		}
		i.Next = nil
		return head
	}

	i := head
	k := head
	j := 1
	for i.Next != nil {
		if j >= n {
			k = k.Next
		}
		j++
		i = i.Next
	}

	k.Val = k.Next.Val
	k.Next = k.Next.Next

	return head
}

func removeNthFromEndOptimized(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	first, second := dummy, dummy
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
