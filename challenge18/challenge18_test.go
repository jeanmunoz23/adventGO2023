package challenge18

import (
	"reflect"
	"testing"
)

func Test_drawClock(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test draw clock",
			args: args{time: "12:45"},
			want: [][]string{
				{" ", " ", "*", "*", "*", "*", " ", " ", " ", "*", " ", "*", "*", "*", "*"},
				{" ", " ", "*", " ", " ", "*", " ", " ", " ", "*", " ", "*", "*", " ", " "},
				{" ", " ", "*", " ", " ", "*", " ", "*", " ", "*", " ", "*", "*", " ", " "},
				{" ", " ", "*", "*", "*", "*", " ", " ", " ", "*", "*", "*", "*", "*", "*"},
				{" ", " ", "*", "*", " ", " ", " ", "*", " ", " ", " ", "*", " ", " ", "*"},
				{" ", " ", "*", "*", " ", " ", " ", " ", " ", " ", "*", "*", " ", " ", "*"},
				{" ", " ", "*", "*", "*", "*", " ", " ", " ", " ", "*", "*", "*", "*", "*"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := drawClock(tt.args.time); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("drawClock() = %v, want %v", got, tt.want)
			}
		})
	}
}
