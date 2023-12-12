package challenge03_test

import (
	"adventgo2023/challenge03"
	"testing"
)

func Test_findNaughtyStep(t *testing.T) {
	type args struct {
		original string
		modified string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test El elfo travieso1 ",
			args: args{"abcd", "abcde"},
			want: "e",
		},
		{
			name: "Test El elfo travieso2 ",
			args: args{"stepfor", "stepor"},
			want: "f",
		},
		{
			name: "Test El elfo travieso3",
			args: args{"abcde", "abcde"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge03.FindNaughtyStep(tt.args.original, tt.args.modified); got != tt.want {
				t.Errorf("findNaughtyStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
