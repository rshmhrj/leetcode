package easy

import (
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "flower,flow",
			args: args{[]string{"flower", "flow"}},
			want: "flow",
		},
		{
			name: "flower,flow,flight",
			args: args{[]string{"flower", "flow", "flight"}},
			want: "fl",
		},
		{
			name: "dog,racecar,car",
			args: args{[]string{"dog", "racecar", "car"}},
			want: "",
		},
		{
			name: "empty list",
			args: args{[]string{""}},
			want: "",
		},
		{
			name: "ab,a",
			args: args{[]string{"ab", "a"}},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fmt.Printf("flower,flow: %v\n", longestCommonPrefix([]string{"flower","flow"})) // "fl"
//fmt.Printf("flower,flow,flight: %v\n", longestCommonPrefix([]string{"flower","flow","flight"})) // "fl"
//fmt.Printf("dog,racecar,car: %v\n", longestCommonPrefix([]string{"dog","racecar","car"})) // ""
