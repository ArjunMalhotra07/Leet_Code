package trees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return inOrderTraversalHelperFunction(root, ans)

}

func inOrderTraversalHelperFunction(root *TreeNode, ans []int) []int {
	if root.Left != nil {
		ans = inOrderTraversalHelperFunction(root.Left, ans)
	}
	if root != nil {
		ans = append(ans, root.Val)
		fmt.Println(ans)
	}

	if root.Right != nil {
		ans = inOrderTraversalHelperFunction(root.Right, ans)
	}
	return ans
}
