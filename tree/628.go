package tree

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/628/
//Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//3
/// \
//9  20
///  \
//15   7
//return its level order traversal as:
//[
//[3],
//[9,20],
//[15,7]
//]
func levelOrder(root *TreeNode) [][]int {
	return traverseNodesWithLevel(root, 0, make([][]int, 0, 100))
}

func traverseNodesWithLevel(root *TreeNode, level int, traversed [][]int) [][]int {
	if root == nil {
		return traversed
	}

	if len(traversed) - 1 >= level {
		traversed[level] = append(traversed[level], root.Val)
	} else {
		traversed = append(traversed, []int{root.Val})
	}
	traversed = traverseNodesWithLevel(root.Left, level + 1, traversed)
	traversed = traverseNodesWithLevel(root.Right, level + 1, traversed)
	return traversed
}

