package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var res []int

	if root == nil {
		return res
	}

	// visit root
	res = append(res, root.Val)
	// visit left subtree
	if root.Left != nil {
		res = append(res, preorderTraversal(root.Left)...)
	}
	// visit right subtree
	if root.Right != nil {
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}
