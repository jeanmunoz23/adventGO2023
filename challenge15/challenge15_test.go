package challenge15

import (
	"reflect"
	"testing"
)

func Test_autonomousDriveOp2(t *testing.T) {
	type args struct {
		store     []string
		movements []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test ['.......',  '...*!*.']",
			args: args{[]string{"..!....", "...*.*."}, []string{"R", "R", "D", "L"}},
			want: []string{".......", "...*!*."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := autonomousDrive(tt.args.store, tt.args.movements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("autonomousDrive() = %v, want %v", got, tt.want)
			}
		})
	}
}
