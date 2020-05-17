package tree

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/
//Given a binary tree, find its maximum depth.
//
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//Note: A leaf is a node with no children.
//
//Example:
//
//Given binary tree [3,9,20,null,null,15,7],
//
//3
/// \
//9  20
///  \
//15   7
//return its depth = 3.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return traverseMaxDepth(root, 0)
}

func traverseMaxDepth(root *TreeNode, maxDep int) int {
	if root == nil {
		return maxDep
	}

	left := traverseMaxDepth(root.Left, maxDep + 1)
	right := traverseMaxDepth(root.Right, maxDep + 1)
	if left > right {
		return left
	}
	return right
}
