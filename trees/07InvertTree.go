package trees

func InvertTree(root *TreeNode) *TreeNode {
	InvertHelper(root)
	return root
}
func InvertHelper(root *TreeNode) {
	if root == nil {
		return
	}
	InvertHelper(root.Left)
	InvertHelper(root.Right)
	var temp *TreeNode = root.Left
	root.Left = root.Right
	root.Right = temp

}
