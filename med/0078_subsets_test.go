package med

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "empty",
			args: args{[]int{}},
			want: [][]int{[]int{}},
		},
		{
			name: "single",
			args: args{[]int{1}},
			want: [][]int{[]int{}, []int{1}},
		},
		{
			name: "double",
			args: args{[]int{1, 2}},
			want: [][]int{[]int{1}, []int{2, 1}, []int{}, []int{2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
