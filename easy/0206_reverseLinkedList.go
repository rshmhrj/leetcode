package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, next, tmp *ListNode

	for tmp = head; tmp != nil; tmp = next {
		next = tmp.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}
