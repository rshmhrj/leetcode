package easy

import (
	"testing"
)

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sqrt(4)=2",
			args: args{4},
			want: 2,
		},
		{
			name: "sqrt(8)=2",
			args: args{8},
			want: 2,
		},
		{
			name: "sqrt(12345)=111",
			args: args{12345},
			want: 111,
		},
		{
			name: "sqrt(654353)=111",
			args: args{654353},
			want: 808,
		},
		{
			name: "sqrt(345903498)=18598",
			args: args{345903498},
			want: 18598,
		},
		{
			name: "sqrt(0)=0",
			args: args{0},
			want: 0,
		},
		{
			name: "sqrt(1)=1",
			args: args{1},
			want: 1,
		},
		{
			name: "sqrt(2)=1",
			args: args{2},
			want: 1,
		},
		{
			name: "sqrt(7)=2",
			args: args{7},
			want: 2,
		},
		{
			name: "sqrt(5)=2",
			args: args{5},
			want: 2,
		},
		{
			name: "sqrt(530698624)=23036",
			args: args{530698624},
			want: 23036,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt(%v)=%v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
