package easy

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "hello",
			args: args{[]byte{'h', 'e', 'l', 'l', 'o'}},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "Hannah",
			args: args{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name: "_; aplanaman!",
			args: args{[]byte{'_', ';', ' ', 'a', 'p', 'l', 'a', 'n', 'a', 'm', 'a', 'n', '!'}},
			want: []byte{'!', 'n', 'a', 'm', 'a', 'n', 'a', 'l', 'p', 'a', ' ', ';', '_'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if string(tt.args.s) != string(tt.want) {
				t.Errorf("reverse(%v)=%v, want %v", tt.name, tt.args.s, tt.want)
			}
		})
	}
}
