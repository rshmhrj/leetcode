package easy

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list1 := makeIntLinkedList([]int{1, 2, 3, 4, 5})
	head1 := list1
	tail1 := list1.Next.Next.Next.Next

	list2 := makeIntLinkedList([]int{1, 2})
	head2 := list2
	tail2 := list2.Next

	list3 := makeIntLinkedList([]int{1, 5, 8, 4, 8, 0, 65, 20, 10, 40, -5000, 5000, 499, 4999, -4999})
	head3 := list3
	tail3 := list3.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "[12345]=[54321]",
			args: args{head1},
			want: tail1,
		},
		{
			name: "[12]=[21]",
			args: args{head2},
			want: tail2,
		},
		{
			name: "[1,5,8,4,8,0,65,20,10,40,-5000,5000,499,4999,-4999]=[-4999,4999,499,5000,-5000,40,10,20,65,0,8,4,8,5,1]",
			args: args{head3},
			want: tail3,
		},
		{
			name: "[]=[]",
			args: args{nil},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
