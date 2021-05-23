package easy

import (
	"reflect"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,null,2,3]=[3,2,1]",
			args: args{makeIntTree([]int{1, -999, 2, -999, -999, 3})},
			want: []int{3, 2, 1},
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
			name: "[1,2,3,4,5]=[4,5,2,3,1]",
			args: args{makeIntTree([]int{1, 2, 3, 4, 5})},
			want: []int{4, 5, 2, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
