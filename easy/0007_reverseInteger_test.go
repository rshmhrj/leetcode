package easy

import (
	"testing"
)

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{

		{
			name: "123",
			args: args{123},
			want: 321,
		},
		{
			name: "-123",
			args: args{-123},
			want: -321,
		},
		{
			name: "120",
			args: args{120},
			want: 21,
		},
		{
			name: "0",
			args: args{0},
			want: 0,
		},
		{
			name: "123456789",
			args: args{123456789},
			want: 987654321,
		},
		{
			name: "100",
			args: args{100},
			want: 1,
		},
		{
			name: "upper bound",
			args: args{1534236469},
			want: 0,
		},
		{
			name: "lower bound",
			args: args{-2147483648},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
