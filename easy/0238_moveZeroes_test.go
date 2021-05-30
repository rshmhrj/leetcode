package easy

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "empty array",
			args: args{[]int{}},
			want: []int{},
		},
		{
			name: "[0]=[0]",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "[1]=[1]",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "[1,0]=[1,0]",
			args: args{[]int{1, 0}},
			want: []int{1, 0},
		},
		{
			name: "[0,1]=[1,0]",
			args: args{[]int{0, 1}},
			want: []int{1, 0},
		},
		{
			name: "[0,1,0,3,12]=[1,3,12,0,0]",
			args: args{[]int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "[0,0,0,0]=[0,0,0,0]",
			args: args{[]int{0, 0, 0, 0}},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "[0,0,0,1]=[1,0,0,0]",
			args: args{[]int{0, 0, 0, 1}},
			want: []int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if moveZeroes(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
