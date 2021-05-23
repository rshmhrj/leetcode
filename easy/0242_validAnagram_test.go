package easy

import (
	"testing"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "anagram/''",
			args: args{"anagram", ""},
			want: false,
		},
		{
			name: "hello/hey",
			args: args{"hello", "hey"},
			want: false,
		},
		{
			name: "cat/car",
			args: args{"cat", "car"},
			want: false,
		},
		{
			name: "anagram/nagaram",
			args: args{"anagram", "nagaram"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
