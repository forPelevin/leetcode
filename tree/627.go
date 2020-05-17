package tree

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/627/
//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//1
/// \
//2   2
/// \ / \
//3  4 4  3
//
//
//But the following [1,2,2,null,3,null,3] is not:
//
//1
/// \
//2   2
//\   \
//3    3
//
//
//Follow up: Solve it both recursively and iteratively.
func isSymmetric(root *TreeNode) bool {
	q := make([]*TreeNode, 0, 100)
	q = append(q, root, root)
	for len(q) > 0 {
		t1 := q[len(q)-1]
		t2 := q[len(q)-2]
		q = q[:len(q)-2]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		q = append(q, t1.Left, t2.Right, t1.Right, t2.Left)
	}
	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	return isSymmetricTrees(root, root)
}

func isSymmetricTrees(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return isSymmetricTrees(t1.Left, t2.Right) && isSymmetricTrees(t1.Right, t2.Left)
}