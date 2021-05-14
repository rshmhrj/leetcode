package livePractice

import (
	"testing"
)

func Test_makeNumberBasic(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "permute 2 = 2",
			args: args{2},
			want: 2,
		},
		{
			name: "permute 3 = 4",
			args: args{3},
			want: 4,
		},
		{
			name: "permute 5 = 1",
			args: args{5},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeNumberBasic(tt.args.x); got != tt.want {
				t.Errorf("makeNumberBasic() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_sum(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "12345=15",
			args: args{12345},
			want: 15,
		},
		{
			name: "20000",
			args: args{20000},
			want: 2,
		},
		{
			name: "99999",
			args: args{99999},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.x); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_stripZeros(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "101",
			args: args{101},
			want: 11,
		},
		{
			name: "20020",
			args: args{20020},
			want: 22,
		},
		{
			name: "1234920002340202340",
			args: args{1234920002340202340},
			want: 1234922342234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stripZeros(tt.args.x); got != tt.want {
				t.Errorf("stripZeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
