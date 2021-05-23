package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	// visit left subtree
	if root.Left != nil {
		res = append(res, inorderTraversal(root.Left)...)
	}
	// visit root
	res = append(res, root.Val)
	// visit right subtree
	if root.Right != nil {
		res = append(res, inorderTraversal(root.Right)...)
	}

	return res
}
