package med

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "blank",
			args: args{""},
			want: 0,
		},
		{
			name: "single space",
			args: args{" "},
			want: 1,
		},
		{
			name: "abcabcbb",
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "bbbbb",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "pwwkew",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "dvdf",
			args: args{"dvdf"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
