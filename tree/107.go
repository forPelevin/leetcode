package tree

//https://leetcode.com/problems/binary-tree-level-order-traversal-ii/submissions/
//107. Binary Tree Level Order Traversal II
//Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//3
/// \
//9  20
///  \
//15   7
//return its bottom-up level order traversal as:
//[
//[15,7],
//[9,20],
//[3]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		qSize := len(q)
		currLvl := make([]int, qSize)
		for i := 0; i < qSize; i++ {
			node := q[i]
			currLvl[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[qSize:]
		res = append([][]int{currLvl}, res...)
	}
	return res
}
