package easy

import (
	"reflect"
	"testing"
)

func Test_fizzBuzz(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "n=3->[1,2,Fizz]",
			args: args{3},
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "n=5->[1,2,Fizz,4,Buzz]",
			args: args{5},
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "n=15->[1,2,Fizz,4,Buzz,6,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz]",
			args: args{15},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
