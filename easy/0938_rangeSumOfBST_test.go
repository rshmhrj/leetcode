package easy

import (
	"testing"
)

func Test_rangeSumBST(t *testing.T) {
	type args struct {
		n    *TreeNode
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[10,5,15,3,7,null,18], low = 7, high = 15 -> 32",
			args: args{makeIntTree([]int{10, 5, 15, 3, 7, -999, 18}), 7, 15},
			want: 32,
		},
		{
			name: "[10,5,15,3,7,13,18,1,null,6], low = 6, high = 10 -> 23",
			args: args{makeIntTree([]int{10, 5, 15, 3, 7, 13, 18, 1, -999, 6}), 6, 10},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeSumBST(tt.args.n, tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("rangeSumBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
