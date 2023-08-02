package trees

func PreorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return PreOrderTraversalHelperFunction(root, ans)

}

func PreOrderTraversalHelperFunction(root *TreeNode, ans []int) []int {
	if root != nil {
		ans = append(ans, root.Val)
	}
	if root.Left != nil {
		ans = PreOrderTraversalHelperFunction(root.Left, ans)
	}

	if root.Right != nil {
		ans = PreOrderTraversalHelperFunction(root.Right, ans)
	}
	return ans
}


