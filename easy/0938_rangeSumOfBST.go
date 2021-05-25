package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(n *TreeNode, low int, high int) int {
	var sum int
	if n.Val >= low && n.Val <= high {
		sum += n.Val
	}

	if n.Left != nil && n.Val >= low {
		sum += rangeSumBST(n.Left, low, high)
	}
	if n.Right != nil && n.Val <= high {
		sum += rangeSumBST(n.Right, low, high)
	}

	return sum
}
