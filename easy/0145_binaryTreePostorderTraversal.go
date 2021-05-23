package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	// visit left subtree
	if root.Left != nil {
		res = append(res, postorderTraversal(root.Left)...)
	}
	// visit right subtree
	if root.Right != nil {
		res = append(res, postorderTraversal(root.Right)...)
	}
	// visit root
	res = append(res, root.Val)
	return res
}
