package easy

import (
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "space array len 1",
			args: args{" "},
			want: 0,
		},
		{
			name: "Hello World",
			args: args{"Hello World"},
			want: 5,
		},
		{
			name: "                            ",
			args: args{"                            "},
			want: 0,
		},
		{
			name: "                            a",
			args: args{"                            a"},
			want: 1,
		},
		{
			name: "                            a ",
			args: args{"                            a "},
			want: 1,
		},
		{
			name: "asdf asdf asdf asdfasdflasdfkjasdflaosiufjewjnflskdnaoisudjfoewnasdnflasdfoisdjfodsafiaosdifnasodfiasoiuroeiansdofjasidfjaosdifjoewinrfalsdkvnaodifajsodfijsdofnaeoirs",
			args: args{"asdf asdf asdf asdfasdflasdfkjasdflaosiufjewjnflskdnaoisudjfoewnasdnflasdfoisdjfodsafiaosdifnasodfiasoiuroeiansdofjasidfjaosdifjoewinrfalsdkvnaodifajsodfijsdofnaeoirs"},
			want: 151,
		},
		{
			name: "asdf asdf asdf asdfasdflasdfkjasdflaosiufjewjnflskdnaoisudjfoewnasdnflasdfoisdjfodsafiaosdifnasodfiasoiuroeiansdofjasidfjaosdifjoewinrfalsdkvnaodifajsodfijsdofnaeoirs ",
			args: args{"asdf asdf asdf asdfasdflasdfkjasdflaosiufjewjnflskdnaoisudjfoewnasdnflasdfoisdjfodsafiaosdifnasodfiasoiuroeiansdofjasidfjaosdifjoewinrfalsdkvnaodifajsodfijsdofnaeoirs "},
			want: 151,
		},
		{
			name: " asdf asdf asdfasdflasdfkjasdflaosiufjewjnflskdnaoisudjfoewnasdnflasdfoisdjfodsafiaosdifnasodfiasoiuroeiansdofjasidfjaosdifjoewinrfalsdkvnaodifajsodfijsdofnaeoirs a",
			args: args{" asdf asdf asdfasdflasdfkjasdflaosiufjewjnflskdnaoisudjfoewnasdnflasdfoisdjfodsafiaosdifnasodfiasoiuroeiansdofjasidfjaosdifjoewinrfalsdkvnaodifajsodfijsdofnaeoirs a"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
