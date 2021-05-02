package easy

import (
	"testing"
)

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,3,5,6], 5",
			args: args{[]int{1, 3, 5, 6}, 5},
			want: 2,
		},
		{
			name: "[1,3,5,6], 2",
			args: args{[]int{1, 3, 5, 6}, 2},
			want: 1,
		},
		{
			name: "[1,3,5,6], 7",
			args: args{[]int{1, 3, 5, 6}, 7},
			want: 4,
		},
		{
			name: "[1,3,5,6], 0",
			args: args{[]int{1, 3, 5, 6}, 0},
			want: 0,
		},
		{
			name: "[1], 0",
			args: args{[]int{1}, 0},
			want: 0,
		},
		{
			name: "[-256,-1,1,256], -256",
			args: args{[]int{-256, -1, 1, 256}, -256},
			want: 0,
		},
		{
			name: "[-256,-1,1,256], -244",
			args: args{[]int{-256, -1, 1, 256}, -244},
			want: 1,
		},
		{
			name: "[-256,-1,1,256], -3",
			args: args{[]int{-256, -1, 1, 256}, -3},
			want: 1,
		},
		{
			name: "[-256,-1,1,256], 2",
			args: args{[]int{-256, -1, 1, 256}, 2},
			want: 3,
		},
		{
			name: "[-256,-1,1,256], 50",
			args: args{[]int{-256, -1, 1, 256}, 50},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
