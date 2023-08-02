package trees

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MinDepth(root.Left)
	rightDepth := MinDepth(root.Right)
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
