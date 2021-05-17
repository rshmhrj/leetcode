package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// deleteNode deletes current node in LinkedList by copying values from next node to this node and updating pointers to skip the next node
func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
	next.Next = nil
}
