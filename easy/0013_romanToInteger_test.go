package easy

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "III",
			args: args{"III"},
			want: 3,
		},
		{
			name: "IV",
			args: args{"IV"},
			want: 4,
		},
		{
			name: "IX",
			args: args{"IX"},
			want: 9,
		},
		{
			name: "LVIII",
			args: args{"LVIII"},
			want: 58,
		},
		{
			name: "MCDLXXVI",
			args: args{"MCDLXXVI"},
			want: 1476,
		},
		{
			name: "MCMXCIV",
			args: args{"MCMXCIV"},
			want: 1994,
		},
		{
			name: "CXCIX",
			args: args{"CXCIX"},
			want: 199,
		},
		{
			name: "CCCXCIX",
			args: args{"CCCXCIX"},
			want: 399,
		},
		{
			name: "MMCCCXCIX",
			args: args{"MMCCCXCIX"},
			want: 2399,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
