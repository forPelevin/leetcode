package linkedlist

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/560/
//Reverse a singly linked list.
//
//Example:
//
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL
//Follow up:
//
//A linked list can be reversed either iteratively or recursively. Could you implement both?

func reverseListIterative(head *ListNode) *ListNode {
	var reversed *ListNode
	i := head
	for i != nil {
		reversed = &ListNode{
			Val:  i.Val,
			Next: reversed,
		}
		i = i.Next
	}
	return reversed
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}
