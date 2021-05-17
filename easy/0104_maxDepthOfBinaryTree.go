package easy

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepth finds depth of binary tree using a Divide and Conquer, Depth-First approach
func maxDepth(root *TreeNode) int {
	// handle empty tree by returning 0 (count of no levels)
	if root == nil {
		return 0
	}
	// handle base case with single root by returning 1 (count of this level)
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// initialize left / right counters
	var l, r int

	// divide and find depth of children, depth first on left branch
	if root.Left != nil {
		l = maxDepth(root.Left)
	}
	if root.Right != nil {
		r = maxDepth(root.Right)
	}

	// conquer by getting max of children plus 1 for this level
	return getMax(l, r) + 1
}

func getMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
