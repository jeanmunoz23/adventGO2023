package challenge10_test

import (
	"adventgo2023/challenge10"
	"fmt"
	"testing"
)

func Test_createChristmasTree(t *testing.T) {
	type args struct {
		ornaments string
		height    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test anna",
			args: args{"123", 4},
			want: `   1
  2 3
 1 2 3
1 2 3 1
   |`,
		},
		{
			name: "Test abab",
			args: args{"*@o", 3},
			want: `  *
 @ o
* @ o
  |`,
		},
		{
			name: "Test abab",
			args: args{"xo", 4},
			want: `   x
  o x
 o x o
x o x o
   |`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge10.CreateChristmasTree(tt.args.ornaments, tt.args.height); got != tt.want {
				t.Errorf("CreateChristmasTree() =\n %v, want\n %v", got, tt.want)
				fmt.Print(got)
			}
		})
	}
}
