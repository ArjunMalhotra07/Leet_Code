package trees

import "math"

func CheckBalancedTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := CheckBalancedTree(root.Left)
	if leftDepth == -1 {
		return -1
	}
	rightDepth := CheckBalancedTree(root.Right)
	if rightDepth == -1 {
		return -1
	}

	diff := math.Abs(float64(leftDepth) - float64(rightDepth))
	if diff > 1 {
		return -1
	}
	return int(math.Max(float64(leftDepth), float64(rightDepth))) + 1
}
func IsBalanced(root *TreeNode) bool {
	return CheckBalancedTree(root) == -1
}
