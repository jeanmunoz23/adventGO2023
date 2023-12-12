package challenge06_test

import (
	"adventgo2023/challenge06"
	"testing"
)

func Test_maxDistance(t *testing.T) {
	type args struct {
		movements string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test movements := >>*<",
			args: args{">>*<"},
			want: 2,
		},
		{
			name: "Test movements := <<<>",
			args: args{"<<<>"},
			want: 2,
		}, {
			name: "Test movements := >***>",
			args: args{">***>"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge06.MaxDistance(tt.args.movements); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
