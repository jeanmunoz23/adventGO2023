package challenge02_test

import (
	"adventgo2023/challenge02"
	"reflect"
	"testing"
)

func Test_manufacture(t *testing.T) {
	type args struct {
		gifts     []string
		materials string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// Ponemos en marcha la fábrica
		// ["tren", "oso"]
		// "tren" SÍ porque sus letras están en "tronesa"
		// "oso" SÍ porque sus letras están en "tronesa"
		// "pelota" NO porque sus letras NO están en "tronesa"
		{
			name: "Ponemos en marcha la fábrica 1 ",
			args: args{[]string{"tren", "oso", "pelota"}, "tronesa"},
			want: []string{"tren", "oso"},
		},
		{
			name: "Ponemos en marcha la fábrica 2 ",
			args: args{[]string{"juego", "puzzle"}, "jlepuz"},
			want: []string{"puzzle"},
		},
		{
			name: "Ponemos en marcha la fábrica 3 ",
			args: args{[]string{"libro", "ps5"}, "psli"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge02.Manufacture(tt.args.gifts, tt.args.materials); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manufacture() = %v, want %v", got, tt.want)
			}
		})
	}
}
