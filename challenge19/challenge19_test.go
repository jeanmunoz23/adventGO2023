package challenge19

import (
	"reflect"
	"testing"
)

func Test_revealSabotage(t *testing.T) {
	type args struct {
		store [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test reval sabotage",
			args: args{store: [][]string{
				{"*", " ", " ", " "},
				{" ", " ", "*", " "},
				{" ", " ", " ", " "},
				{"*", " ", " ", " "},
			},
			},
			want: [][]string{
				{"*", "2", "1", "1"},
				{"1", "2", "*", "1"},
				{"1", "2", "1", "1"},
				{"*", "1", " ", " "},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revealSabotage(tt.args.store); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revealSabotage() = %v, want %v", got, tt.want)
			}
		})
	}
}
