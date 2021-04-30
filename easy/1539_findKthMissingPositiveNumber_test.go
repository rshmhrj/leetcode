package easy

import (
	"testing"
)

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2,3,4,7,11",
			args: args{arr: []int{2, 3, 4, 7, 11}, k: 5},
			want: 9, // 1,5,6,8,9
		},
		{
			name: "1,2,3,4",
			args: args{arr: []int{1, 2, 3, 4}, k: 2},
			want: 6, // 5,6
		},
		{
			name: "5,6,7,8,9",
			args: args{arr: []int{5, 6, 7, 8, 9}, k: 9},
			want: 14, // 1,2,3,4,10,11,12,13,14
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
