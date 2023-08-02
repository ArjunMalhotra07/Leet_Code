package trees

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return postOrderTraversalHelperFunction(root, ans)

}

func postOrderTraversalHelperFunction(root *TreeNode, ans []int) []int {

	if root.Left != nil {
		ans = postOrderTraversalHelperFunction(root.Left, ans)
	}

	if root.Right != nil {
		ans = postOrderTraversalHelperFunction(root.Right, ans)
	}
	if root != nil {
		ans = append(ans, root.Val)
	}
	return ans
}


