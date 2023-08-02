package trees

func PostorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return PostOrderTraversalHelperFunction(root, ans)

}

func PostOrderTraversalHelperFunction(root *TreeNode, ans []int) []int {

	if root.Left != nil {
		ans = PostOrderTraversalHelperFunction(root.Left, ans)
	}

	if root.Right != nil {
		ans = PostOrderTraversalHelperFunction(root.Right, ans)
	}
	if root != nil {
		ans = append(ans, root.Val)
	}
	return ans
}


