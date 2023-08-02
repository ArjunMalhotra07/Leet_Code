package trees

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	if leftDepth > rightDepth {
		return 1 + leftDepth
	} else {
		return 1 + rightDepth
	}
}
