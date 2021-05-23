package easy

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,null,2,3]=[1,3,2]",
			args: args{makeIntTree([]int{1, -999, 2, -999, -999, 3})},
			want: []int{1, 3, 2},
		},
		{
			name: "[]=[]",
			args: args{nil},
			want: nil,
		},
		{
			name: "[1]=[1]",
			args: args{makeIntTree([]int{1})},
			want: []int{1},
		},
		{
			name: "[1,2,3,4,5]=[4,2,5,1,3]",
			args: args{makeIntTree([]int{1, 2, 3, 4, 5})},
			want: []int{4, 2, 5, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); (len(got) != len(tt.want)) || (!reflect.DeepEqual(got, tt.want)) {
				t.Errorf("inorderTraversal() = %v len(%v), want %v len(%v)", got, len(got), tt.want, len(tt.want))
			}
		})
	}
}

func makeIntTree(a []int) *TreeNode {
	var tmp []*TreeNode
	if len(a) == 0 {
		return &TreeNode{}
	}
	for i, v := range a {
		if v == -999 {
			tmp = append(tmp, nil)
			continue
		}
		node := &TreeNode{
			Val:   v,
			Left:  nil,
			Right: nil,
		}
		tmp = append(tmp, node)
		if i == 0 {
			continue
		}

		if i%2 == 0 {
			tmp[(i-1)/2].Right = node
		}
		if i%2 != 0 {
			tmp[(i-1)/2].Left = node
		}
	}
	return tmp[0]
}
