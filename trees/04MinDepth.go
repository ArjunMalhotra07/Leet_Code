package trees

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth < rightDepth {
		if leftDepth != 0 {
			return 1 + leftDepth
		} else {
			return 1 + rightDepth
		}

	} else {
		if rightDepth != 0 {
			return 1 + rightDepth
		} else {
			return 1 + leftDepth
		}

	}
}
