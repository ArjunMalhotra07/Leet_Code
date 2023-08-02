package trees

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return InOrderTraversalHelperFunction(root, ans)

}

func InOrderTraversalHelperFunction(root *TreeNode, ans []int) []int {
	if root.Left != nil {
		ans = InOrderTraversalHelperFunction(root.Left, ans)
	}
	if root != nil {
		ans = append(ans, root.Val)
		fmt.Println(ans)
	}

	if root.Right != nil {
		ans = InOrderTraversalHelperFunction(root.Right, ans)
	}
	return ans
}
