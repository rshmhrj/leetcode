package easy

import (
	"testing"
)

func Test_buddyStrings(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "abcd & cbad",
			args: args{a: "abcd", b: "cbad"},
			want: true,
		},
		{
			name: "ab & ab",
			args: args{a: "ab", b: "ab"},
			want: false,
		},
		{
			name: "aa & aa",
			args: args{a: "aa", b: "aa"},
			want: true,
		},
		{
			name: "aaaaaaabc & aaaaaaacb",
			args: args{a: "aaaaaaabc", b: "aaaaaaacb"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
