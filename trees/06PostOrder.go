/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	return traversalHelper(root, ans)

}

func traversalHelper(root *TreeNode, ans []int) []int {

	if root.Left != nil {
		ans = traversalHelper(root.Left, ans)
	}

	if root.Right != nil {
		ans = traversalHelper(root.Right, ans)
	}
	if root != nil {
		ans = append(ans, root.Val)
	}
	return ans
}


