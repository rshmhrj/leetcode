package easy

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1 := &List{}
	list1.append(&ListNode{Val: 1})
	list1.append(&ListNode{Val: 2})
	list1.append(&ListNode{Val: 4})

	list2 := &List{}
	list2.append(&ListNode{Val: 1})
	list2.append(&ListNode{Val: 3})
	list2.append(&ListNode{Val: 4})

	list3 := &List{}
	list3.append(list1.head)
	list3.append(list2.head)
	list3.append(list1.head.Next)
	list3.append(list2.head.Next)
	list3.append(list1.head.Next.Next)
	list3.append(list2.head.Next.Next)

	list4 := &List{}
	list5 := &List{}
	list6 := &List{}
	list7 := &List{}
	list7.append(&ListNode{Val: 0})
	list8 := &List{}
	list8.append(&ListNode{Val: -9})
	list8.append(&ListNode{Val: 3})
	list9 := &List{}
	list9.append(&ListNode{Val: 5})
	list9.append(&ListNode{Val: 7})
	list10 := &List{}
	list10.append(&ListNode{Val: -9})
	list10.append(&ListNode{Val: 3})
	list10.append(&ListNode{Val: 5})
	list10.append(&ListNode{Val: 7})

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "124 and 134",
			args: args{l1: list1.head, l2: list2.head},
			want: list3.head,
		},
		{
			name: "134 and 124",
			args: args{l1: list2.head, l2: list1.head},
			want: list3.head,
		},
		{
			name: "[] and []",
			args: args{l1: list4.head, l2: list5.head},
			want: list6.head,
		},
		{
			name: "[] and [0]",
			args: args{l1: list4.head, l2: list7.head},
			want: list7.head,
		},
		{
			name: "[0] and []",
			args: args{l1: list7.head, l2: list4.head},
			want: list7.head,
		},
		{
			name: "[-9,3] and [5,7]",
			args: args{l1: list8.head, l2: list9.head},
			want: list10.head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
