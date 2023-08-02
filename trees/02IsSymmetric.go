package trees

func IsSymmetric(root *TreeNode) bool {
	return IsSymmetricHelper(root.Left, root.Right)
}

func IsSymmetricHelper(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return IsSymmetricHelper(left.Left, right.Right) && IsSymmetricHelper(left.Right, right.Left)
}

