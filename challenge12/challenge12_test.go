package challenge12_test

import (
	"adventgo2023/challenge12"
	"testing"
)

func TestCheckIsValidCopy(t *testing.T) {
	type args struct {
		original string
		copy     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Returns true",
			args: args{"Santa Claus is coming", "sa#ta cl#us is comin#"},
			want: true,
		},
		{
			name: "Test Returns false: por la p inicial",
			args: args{"Santa Claus is coming", "p#nt: cla#s #s c+min#"},
			want: false,
		},
		{
			name: "Test Returns true: s#+:. c:. s",
			args: args{"Santa Claus", "s#+:. c:. s"},
			want: true,
		},
		{
			name: "Test Returns false: (hay un # donde no debería)",
			args: args{"Santa Claus", "s#+:.#c:. s"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge12.CheckIsValidCopy(tt.args.original, tt.args.copy); got != tt.want {
				t.Errorf("CheckIsValidCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
