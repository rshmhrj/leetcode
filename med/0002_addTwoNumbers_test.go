package med

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := NewList()
	l1.push(3)
	l1.push(4)
	l1.push(2)

	l2 := NewList()
	l2.push(4)
	l2.push(6)
	l2.push(5)

	l3 := NewList()
	l3.push(8)
	l3.push(0)
	l3.push(7)

	l4 := NewList()
	l4.push(9)
	l4.push(9)
	l4.push(9)
	l4.push(9)
	l4.push(9)
	l4.push(9)
	l4.push(9)

	l5 := NewList()
	l5.push(9)
	l5.push(9)
	l5.push(9)
	l5.push(9)

	l6 := NewList()
	l6.push(1)
	l6.push(0)
	l6.push(0)
	l6.push(0)
	l6.push(9)
	l6.push(9)
	l6.push(9)
	l6.push(8)

	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "342 + 465 = 807",
			args: args{l1: l1.head, l2: l2.head},
			want: l3.head.Val,
		},
		{
			name: "9999999 + 9999 = 10009998",
			args: args{l1: l4.head, l2: l5.head},
			want: l6.head.Val,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2).Val; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
