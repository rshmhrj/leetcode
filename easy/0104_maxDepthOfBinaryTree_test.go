package easy

import (
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,9,20,null,null,15,7]=3",
			args: args{makeIntTree([]int{3, 9, 20, -999, -999, 15, 7})},
			want: 3,
		},
		{
			name: "[1,null,2]=2",
			args: args{makeIntTree([]int{1, -999, 2})},
			want: 2,
		},
		{
			name: "[]=0",
			args: args{nil},
			want: 0,
		},
		{
			name: "[0]=1",
			args: args{makeIntTree([]int{0})},
			want: 1,
		},
		{
			name: "[1,2,3,4,5]=[1,2,4,5,3]",
			args: args{makeIntTree([]int{1, 2, 3, 4, 5})},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
