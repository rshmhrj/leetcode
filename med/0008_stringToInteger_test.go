package med

import (
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "42",
			args: args{"42"},
			want: 42,
		},
		{
			name: "no number here",
			args: args{"no number here"},
			want: 0,
		},
		{
			name: "-42",
			args: args{"-42"},
			want: -42,
		},
		{
			name: "    -42",
			args: args{"    -42"},
			want: -42,
		},
		{
			name: "4193 with words",
			args: args{"4193 with words"},
			want: 4193,
		},
		{
			name: "words and 987",
			args: args{"words and 987"},
			want: 0,
		},
		{
			name: "-91283472332",
			args: args{"-91283472332"},
			want: -2147483648,
		},
		{
			name: "-",
			args: args{"-"},
			want: 0,
		},
		{
			name: "+",
			args: args{"+"},
			want: 0,
		},
		{
			name: "+1",
			args: args{"+1"},
			want: 1,
		},
		{
			name: "+1",
			args: args{"+1"},
			want: 1,
		},
		{
			name: "+12",
			args: args{"+12"},
			want: 12,
		},
		{
			name: "+-12",
			args: args{"+-12"},
			want: 0,
		},
		{
			name: "-+12",
			args: args{"-+12"},
			want: 0,
		},
		{
			name: "2",
			args: args{"2"},
			want: 2,
		},
		{
			name: "20000000000000000000",
			args: args{"20000000000000000000"},
			want: 2147483647,
		},
		{
			name: "-000000000000000000000000000000000000000000000000001",
			args: args{"-000000000000000000000000000000000000000000000000001"},
			want: -1,
		},
		{
			name: "  +  413",
			args: args{"  +  413"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
