package easy

import (
	"testing"
)

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[3,2,1]",
			args: args{[]int{3, 2, 1}},
			want: 1,
		},
		{
			name: "[1,2]",
			args: args{[]int{1, 2}},
			want: 2,
		},
		{
			name: "[2,2,3,1]",
			args: args{[]int{2, 2, 3, 1}},
			want: 1,
		},
		{
			name: "[1,1,2]",
			args: args{[]int{1, 1, 2}},
			want: 2,
		},
		{
			name: "[1,2,2]",
			args: args{[]int{1, 2, 2}},
			want: 2,
		},
		{
			name: "[5,2,2]",
			args: args{[]int{5, 2, 2}},
			want: 5,
		},
		{
			name: "[1,-2147483648,2]",
			args: args{[]int{1, -2147483648, 2}},
			want: -2147483648,
		},
		{
			name: "[-1,-2147483648,-2147483648]",
			args: args{[]int{-1, -2147483648, -2147483648}},
			want: -1,
		},
		{
			name: "[5,2,4,1,3,6,0]",
			args: args{[]int{5, 2, 4, 1, 3, 6, 0}},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
