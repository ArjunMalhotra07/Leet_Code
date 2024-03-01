package trees

import "math"

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
