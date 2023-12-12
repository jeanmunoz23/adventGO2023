package challenge01_test

import (
	"adventgo2023/challenge01"
	"testing"
)

func TestFindFirstRepeated(t *testing.T) {
	type args struct {
		gifts []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{ // Aunque el 2 y el 3 se repiten
		// el 3 aparece primero por segunda vez
		{
			name: "¡Primer regalo repetido! es el 3 ",
			args: args{[]int{2, 1, 3, 5, 3, 2}},
			want: 3,
		},
		{
			name: "¡Primer regalo repetido! es el 5",
			args: args{[]int{5, 1, 5, 1}},
			want: 5,
		},
		{
			name: "Es -1 ya que no se repite ningún número ",
			args: args{[]int{1, 2, 3, 4}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge01.FindFirstRepeated(tt.args.gifts); got != tt.want {
				t.Errorf("FindFirstRepeated() = %v, want %v", got, tt.want)
			}
		})
	}
}
