package easy

import (
	"testing"
)

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "0=0",
			args: args{0},
			want: 0,
		},
		{
			name: "1=0",
			args: args{0},
			want: 0,
		},
		{
			name: "2=0",
			args: args{2},
			want: 0,
		},
		{
			name: "3=1",
			args: args{3},
			want: 1,
		},
		{
			name: "5=2",
			args: args{5},
			want: 2,
		},
		{
			name: "10=4",
			args: args{10},
			want: 4,
		},
		{
			name: "25=9",
			args: args{25},
			want: 9,
		},
		{
			name: "100=25",
			args: args{100},
			want: 25,
		},
		{
			name: "10000=1229",
			args: args{10000},
			want: 1229,
		},
		{
			name: "5000000=348513",
			args: args{5000000},
			want: 348513,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
