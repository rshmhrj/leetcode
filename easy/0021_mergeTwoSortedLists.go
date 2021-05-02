package easy

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
	len  int
}

func (l *List) append(n *ListNode) {
	node := &ListNode{}
	node.Val = n.Val
	if l.head == nil {
		l.head = node
		return
	}
	var p *ListNode
	for p = l.head; p.Next != nil; p = p.Next {
		continue
	}
	p.Next = node
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &List{}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	p, q := l1, l2
	for p != nil && q != nil {
		if p.Val <= q.Val {
			l.append(p)
			p = p.Next
			continue
		}
		if q.Val < p.Val {
			l.append(q)
			q = q.Next
			continue
		}
	}
	for p != nil {
		l.append(p)
		p = p.Next
	}
	for q != nil {
		l.append(q)
		q = q.Next
	}
	return l.head
}
