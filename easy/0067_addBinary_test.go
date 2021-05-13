package easy

import (
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "11 + 10 = 101",
			args: args{"11", "10"},
			want: "101",
		},
		{
			name: "11 + 1 = 100",
			args: args{"11", "1"},
			want: "100",
		},
		{
			name: "1 + 11 = 100",
			args: args{"1", "11"},
			want: "100",
		},
		{
			name: "1010 + 1011 = 10101",
			args: args{"1010", "1011"},
			want: "10101",
		},
		{
			name: "0 + 0 = 0",
			args: args{"0", "0"},
			want: "0",
		},
		{
			name: "0 + 1 = 1",
			args: args{"0", "1"},
			want: "1",
		},
		{
			name: "1 + 1 = 10",
			args: args{"1", "1"},
			want: "10",
		},
		{
			name: "1 + 111 = 1000",
			args: args{"1", "111"},
			want: "1000",
		},
		{
			name: "1111111101010101010101111111111110101 + 1111111111111111111111111111111111111111 = 10001111111101010101010101111111111110100",
			args: args{"1111111101010101010101111111111110101", "1111111111111111111111111111111111111111"},
			want: "10001111111101010101010101111111111110100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
