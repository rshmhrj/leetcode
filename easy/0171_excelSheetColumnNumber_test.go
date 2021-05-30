package easy

import (
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "A=1",
			args: args{"A"},
			want: 1,
		},
		{
			name: "Z=26",
			args: args{"Z"},
			want: 26,
		},
		{
			name: "AA=27",
			args: args{"AA"},
			want: 27,
		},
		{
			name: "AB=28",
			args: args{"AB"},
			want: 28,
		},
		{
			name: "AZ=52",
			args: args{"AZ"},
			want: 52,
		},
		{
			name: "BA=53",
			args: args{"BA"},
			want: 53,
		},
		{
			name: "ZY=701",
			args: args{"ZY"},
			want: 701,
		},
		{
			name: "ZZ=702",
			args: args{"ZZ"},
			want: 702,
		},
		{
			name: "AAA=703",
			args: args{"AAA"},
			want: 703,
		},
		{
			name: "ABA=729",
			args: args{"ABA"},
			want: 729,
		},
		{
			name: "AZZ=1378",
			args: args{"AZZ"},
			want: 1378,
		},
		{
			name: "BAA=1379",
			args: args{"BAA"},
			want: 1379,
		},
		{
			name: "BBB=1406",
			args: args{"BBB"},
			want: 1406,
		},
		{
			name: "CCC=2109",
			args: args{"CCC"},
			want: 2109,
		},
		{
			name: "ZZZ=18278",
			args: args{"ZZZ"},
			want: 18278,
		},
		{
			name: "FFFF=109674",
			args: args{"FFFF"},
			want: 109674,
		},
		{
			name: "FXSHRXW=2147483647",
			args: args{"FXSHRXW"},
			want: 2147483647,
		},
		{
			name: "FXSHRXV=2147483646",
			args: args{"FXSHRXV"},
			want: 2147483646,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := titleToNumber(tt.args.columnTitle); got != tt.want {
				t.Errorf("titleToNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
