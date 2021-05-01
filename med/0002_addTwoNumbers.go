package med

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Helper struct for testing
type List struct {
	head *ListNode
	len  int
}

// Helper func to create list for testing
func NewList() *List {
	return &List{}
}

// Helper func to push values to list for testing
func (l *List) push(val int) {
	n := &ListNode{
		Val:  val,
		Next: l.head,
	}
	l.head = n
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	nextCase := determineHasNext(l1.Next, l2.Next)

	response := &ListNode{
		Val:  sum,
		Next: nil,
	}

	if nextCase == noHasNext && sum < 10 {
		return response
	}

	if nextCase == noHasNext && sum >= 10 {
		response.Val = sum % 10
		response.Next = addTwoNumbers(&ListNode{Val: 1}, &ListNode{Val: 0})
		return response
	}

	if nextCase == bothHasNext && sum < 10 {
		response.Next = addTwoNumbers(l1.Next, l2.Next)
		return response
	}

	if nextCase == bothHasNext && sum >= 10 {
		response.Val = sum % 10
		l1.Next.Val += 1
		response.Next = addTwoNumbers(l1.Next, l2.Next)
		return response
	}

	if nextCase == firstHasNext && sum < 10 {
		response.Next = addTwoNumbers(l1.Next, &ListNode{Val: 0})
		return response
	}

	if nextCase == secondHasNext && sum < 10 {
		response.Next = addTwoNumbers(&ListNode{Val: 0}, l2.Next)
		return response
	}

	if nextCase == firstHasNext && sum >= 10 {
		response.Val = sum % 10
		l1.Next.Val += 1
		response.Next = addTwoNumbers(l1.Next, &ListNode{Val: 0})
		return response
	}

	if nextCase == secondHasNext && sum >= 10 {
		response.Val = sum % 10
		l2.Next.Val += 1
		response.Next = addTwoNumbers(&ListNode{Val: 0}, l2.Next)
		return response
	}

	return response
}

const (
	noHasNext     = 0
	firstHasNext  = 1
	secondHasNext = 2
	bothHasNext   = 3
)

func determineHasNext(first *ListNode, second *ListNode) int {
	if first == nil && second == nil {
		return noHasNext
	}
	if first != nil && second == nil {
		return firstHasNext
	}
	if first == nil && second != nil {
		return secondHasNext
	}
	if first != nil && second != nil {
		return bothHasNext
	}
	return 0
}
