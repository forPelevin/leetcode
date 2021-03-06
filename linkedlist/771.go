package linkedlist

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/93/linked-list/771/
//Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	i := res
	for l1 != nil || l2 != nil {
		if l1 == nil {
			i.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			i.Next = l1
			l1 = nil
			continue
		}
		if l1.Val <= l2.Val {
			i.Next = l1
			i = i.Next
			l1 = l1.Next
			continue
		}
		i.Next = l2
		i = i.Next
		l2 = l2.Next
	}
	return res.Next
}
