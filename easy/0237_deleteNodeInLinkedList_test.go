package easy

import (
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	list1 := makeIntLinkedList([]int{4, 5, 1, 9})
	tests := []struct {
		name string
		args args
	}{
		{
			name: "[4,5,1,9]-5=[4,1,9]",
			args: args{list1.Next},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if deleteNode(tt.args.node); list1.Next.Val == 5 {
				t.Errorf("Did not delete 5 in list: %v->%v->%v", list1.Val, list1.Next.Val, list1.Next.Next.Val)
			}
		})
	}
}

func makeIntLinkedList(a []int) *ListNode {
	var tmp []*ListNode

	for i, v := range a {
		node := &ListNode{
			Val:  v,
			Next: nil,
		}
		tmp = append(tmp, node)

		if i == 0 {
			continue
		}

		tmp[i-1].Next = node
	}

	return tmp[0]
}
