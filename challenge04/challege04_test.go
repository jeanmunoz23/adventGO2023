package challenge04_test

import (
	"adventgo2023/challenge04"
	"testing"
)

func Test_decode(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Dale la vuelta a los paréntesis: hola mundo ",
			args: args{"hola (odnum)"},
			want: "hola mundo",
		},
		{
			name: "Test Dale la vuelta a los paréntesis: hello world! ",
			args: args{"(olleh) (dlrow)!"},
			want: "hello world!",
		},
		{
			name: "Test Dale la vuelta a los paréntesis: santaclaus ",
			args: args{"sa(u(cla)atn)s"},
			want: "santaclaus",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge04.Decode(tt.args.message); got != tt.want {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
