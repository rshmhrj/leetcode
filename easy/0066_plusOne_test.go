package easy

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "123",
			args: args{[]int{1, 2, 3}},
			want: []int{1, 2, 4},
		},
		{
			name: "4321",
			args: args{[]int{4, 3, 2, 1}},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "0",
			args: args{[]int{0}},
			want: []int{1},
		},
		{
			name: "9",
			args: args{[]int{9}},
			want: []int{1, 0},
		},
		{
			name: "19",
			args: args{[]int{1, 9}},
			want: []int{2, 0},
		},
		{
			name: "119",
			args: args{[]int{1, 1, 9}},
			want: []int{1, 2, 0},
		},
		{
			name: "99",
			args: args{[]int{9, 9}},
			want: []int{1, 0, 0},
		},
		{
			name: "199",
			args: args{[]int{1, 9, 9}},
			want: []int{2, 0, 0},
		},
		{
			name: "999",
			args: args{[]int{9, 9, 9}},
			want: []int{1, 0, 0, 0},
		},
		{
			name: "9,9,9,9,9,9,9,9,9,9",
			args: args{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
			want: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name: "8,9",
			args: args{[]int{8, 9}},
			want: []int{9, 0},
		},
		{
			name: "9,8,9",
			args: args{[]int{9, 8, 9}},
			want: []int{9, 9, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
