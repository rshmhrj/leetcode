package easy

import (
	"testing"
)

func Test_singleNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[2,2,1]=1",
			args: args{[]int{2, 2, 1}},
			want: 1,
		},
		{
			name: "[4,1,2,1,2]=4",
			args: args{[]int{4, 1, 2, 1, 2}},
			want: 4,
		},
		{
			name: "[1]=1",
			args: args{[]int{1}},
			want: 1,
		},
		{
			name: "[1,2,3,4,4,3,2,1,5,6,7,6,7,8,8,9,5,9,0]=0",
			args: args{[]int{1, 2, 3, 4, 4, 3, 2, 1, 5, 6, 7, 6, 7, 8, 8, 9, 5, 9, 0}},
			want: 0,
		},
		{
			name: "[0,1,2,3,4,4,3,2,1,5,6,7,6,7,8,8,9,5,9]=0",
			args: args{[]int{0, 1, 2, 3, 4, 4, 3, 2, 1, 5, 6, 7, 6, 7, 8, 8, 9, 5, 9}},
			want: 0,
		},
		{
			name: "[1,2,3,4,4,3,2,1,0,5,6,7,6,7,8,8,9,5,9]=0",
			args: args{[]int{1, 2, 3, 4, 4, 3, 2, 1, 0, 5, 6, 7, 6, 7, 8, 8, 9, 5, 9}},
			want: 0,
		},
		{
			name: "[1,2,3,4,4,0,3,2,1,5,6,7,6,7,8,8,9,5,9]=0",
			args: args{[]int{1, 2, 3, 4, 4, 0, 3, 2, 1, 5, 6, 7, 6, 7, 8, 8, 9, 5, 9}},
			want: 0,
		},
		{
			name: "[1,2,3,4,4,3,2,1,5,6,7,6,0,7,8,8,9,5,9]=0",
			args: args{[]int{1, 2, 3, 4, 4, 3, 2, 1, 5, 6, 7, 6, 0, 7, 8, 8, 9, 5, 9}},
			want: 0,
		},
		{
			name: "[6,7,6,0,7]=0",
			args: args{[]int{6, 7, 6, 0, 7}},
			want: 0,
		},
		{
			name: "[1,4,4,0,1]=0",
			args: args{[]int{1, 4, 4, 0, 1}},
			want: 0,
		},
		{
			name: "[-1,-4,-4,0,-1]=0",
			args: args{[]int{-1, -4, -4, 0, -1}},
			want: 0,
		},
		{
			name: "[1,0,1]=0",
			args: args{[]int{1, 0, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
