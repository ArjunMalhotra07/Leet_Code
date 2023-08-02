package trees

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return preOrderTraversalHelperFunction(root, ans)

}

func preOrderTraversalHelperFunction(root *TreeNode, ans []int) []int {
	if root != nil {
		ans = append(ans, root.Val)
	}
	if root.Left != nil {
		ans = preOrderTraversalHelperFunction(root.Left, ans)
	}

	if root.Right != nil {
		ans = preOrderTraversalHelperFunction(root.Right, ans)
	}
	return ans
}


