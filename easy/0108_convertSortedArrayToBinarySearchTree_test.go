package easy

import (
	"reflect"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "[1,3]=[3,1]",
			args: args{[]int{1, 3}},
			want: makeIntTree([]int{3, 1}),
		},
		{
			name: "[-10,-3,0,5,9]=[0,-10,5,-999,-3,-999,9]",
			args: args{[]int{-10, -3, 0, 5, 9}},
			want: makeIntTree([]int{0, -3, 9, -10, -999, 5}),
		},
		{
			name: "[-1234,-234,-200,-100,-50,-25,-10,-5,-1,0,1,2,3,4,10,20,200,250,700,800,1000,2000]=[1,-50,200,-234,-5,4,800,-1234,-100,-10,0,3,20,700,2000,-999,-999,-200,-999,-25,-999,-1,-999,2,-999,10,-999,250,-999,1000]",
			args: args{[]int{-1234, -234, -200, -100, -50, -25, -10, -5, -1, 0, 1, 2, 3, 4, 10, 20, 200, 250, 700, 800, 1000, 2000}},
			want: makeIntTree([]int{2, -25, 250, -200, -1, 10, 1000, -234, -50, -5, 1, 4, 200, 800, 2000, -1234, -999, -100, -999, -10, -999, 0, -999, 3, -999, 20, -999, 700}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedArrayToBST(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
