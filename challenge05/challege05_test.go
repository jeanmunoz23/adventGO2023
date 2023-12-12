package challenge05

import (
	"reflect"
	"testing"
)

func Test_cyberReindeer(t *testing.T) {
	type args struct {
		road string
		time int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "Test Santa 🎅 está probando su nuevo trineo eléctrico, el CyberReindeer ",
			args: args{"S..|...|..", 10},
			want: []string{"S..|...|..", ".S.|...|..", "..S|...|..",
				"..S|...|..", "..S|...|..", "...S...*..", "...*S..*..",
				"...*.S.*..", "...*...S..", "...*....S."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cyberReindeer(tt.args.road, tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cyberReindeer() = %v, want %v", got, tt.want)
			}
		})
	}
}
