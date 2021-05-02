package easy

import (
	"testing"
)

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "empty needle",
			args: args{"asdf", ""},
			want: 0,
		},
		{
			name: "empty haystack",
			args: args{"", "asdf"},
			want: -1,
		},
		{
			name: "empty both",
			args: args{"", ""},
			want: 0,
		},
		{
			name: "bigger needle",
			args: args{"asdf", "asdfasdf"},
			want: -1,
		},
		{
			name: "hello,ll",
			args: args{"hello", "ll"},
			want: 2,
		},
		{
			name: "aaaaa,bba",
			args: args{"aaaaa", "bba"},
			want: -1,
		},
		{
			name: "asdf,asd",
			args: args{"asdf", "asd"},
			want: 0,
		},
		{
			name: "asdf,df",
			args: args{"asdf", "df"},
			want: 2,
		},
		{
			name: "hello there,here",
			args: args{"hello there", "here"},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
