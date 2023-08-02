package trees

func invertTree(root *TreeNode) *TreeNode {
	invertHelper(root)
	return root
}
func invertHelper(root *TreeNode) {
	if root == nil {
		return
	}
	invertHelper(root.Left)
	invertHelper(root.Right)
	var temp *TreeNode
	temp = root.Left
	root.Left = root.Right
	root.Right = temp

}
