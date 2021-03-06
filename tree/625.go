package tree

import "math"

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/625/
//Given a binary tree, determine if it is a valid binary search tree (BST).
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than the node's key.
//The right subtree of a node contains only nodes with keys greater than the node's key.
//Both the left and right subtrees must also be binary search trees.
//
//
//Example 1:
//
//2
/// \
//1   3
//
//Input: [2,1,3]
//Output: true
//Example 2:
//
//5
/// \
//1   4
/// \
//3   6
//
//Input: [5,1,4,null,null,3,6]
//Output: false
//Explanation: The root node's value is 5 but its right child's value is 4.
func isValidBST(root *TreeNode) bool {
	return traverseBST(root, math.MinInt64, math.MaxInt64)
}

func traverseBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val <= min {
		return false
	}
	if node.Val >= max {
		return false
	}
	if !traverseBST(node.Right, node.Val, max) {
		return false
	}
	if !traverseBST(node.Left, min, node.Val) {
		return false
	}

	return true
}
