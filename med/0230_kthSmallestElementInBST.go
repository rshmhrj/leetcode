package med

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res []int

	if root == nil {
		return 0
	}

	res = traverse(root)

	// return kth value
	return res[k-1]
}

func traverse(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	if root.Left != nil {
		res = append(res, traverse(root.Left)...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, traverse(root.Right)...)
	}
	return res
}
