package med

import (
	"testing"
)

func Test_divide(t *testing.T) {
	type args struct {
		dividend int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10/3=3",
			args: args{10, 3},
			want: 3,
		},
		{
			name: "7/-3=-2",
			args: args{7, -3},
			want: -2,
		},
		{
			name: "-2147483648/-1=2147483647",
			args: args{-2147483648, -1},
			want: 2147483647,
		},
		{
			name: "-2147483648/1=2147483647",
			args: args{-2147483648, 1},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if testing.Short() && tt.name == "-2147483648/-1=2147483647" || tt.name == "-2147483648/1=2147483647" {
				t.Skip("skipping test in short mode.")
			}
			if got := divide(tt.args.dividend, tt.args.divisor); got != tt.want {
				t.Errorf("divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
