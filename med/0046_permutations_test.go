package med

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "single",
			args: args{[]int{1}},
			want: [][]int{[]int{1}},
		},
		{
			name: "double",
			args: args{[]int{0, 1}},
			want: [][]int{[]int{0, 1}, []int{1, 0}},
		},
		{
			name: "triple",
			args: args{[]int{1, 2, 3}},
			want: [][]int{[]int{1, 2, 3}, []int{2, 1, 3}, []int{2, 3, 1}, []int{1, 3, 2}, []int{3, 1, 2}, []int{3, 2, 1}},
		},
		{
			name: "quad",
			args: args{[]int{1, 3, 6, 9}},
			want: [][]int{[]int{1, 3, 6, 9}, []int{3, 1, 6, 9}, []int{3, 6, 1, 9}, []int{3, 6, 9, 1}, []int{1, 6, 3, 9}, []int{6, 1, 3, 9}, []int{6, 3, 1, 9}, []int{6, 3, 9, 1}, []int{1, 6, 9, 3}, []int{6, 1, 9, 3}, []int{6, 9, 1, 3}, []int{6, 9, 3, 1}, []int{1, 3, 9, 6}, []int{3, 1, 9, 6}, []int{3, 9, 1, 6}, []int{3, 9, 6, 1}, []int{1, 9, 3, 6}, []int{9, 1, 3, 6}, []int{9, 3, 1, 6}, []int{9, 3, 6, 1}, []int{1, 9, 6, 3}, []int{9, 1, 6, 3}, []int{9, 6, 1, 3}, []int{9, 6, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
