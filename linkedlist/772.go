package linkedlist

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/772/
//Given a singly linked list, determine if it is a palindrome.
//
//Example 1:
//
//Input: 1->2
//Output: false
//Example 2:
//
//Input: 1->2->2->1
//Output: true
//Follow up:
//Could you do it in O(n) time and O(1) space?

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	slow = reverse(slow)
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
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
