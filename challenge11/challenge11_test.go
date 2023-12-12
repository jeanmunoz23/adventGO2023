package challenge11_test

import (
	"adventgo2023/challenge11"
	"testing"
)

func Test_getIndexsForPalindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test anna",
			args: args{"anna"},
			want: "[]",
		},
		{
			name: "Test abab",
			args: args{"abab"},
			want: "[0,1]",
		},
		{
			name: "Test abac",
			args: args{"abac"},
			want: "null",
		},
		{
			name: "Test aaaaaaaa",
			args: args{"aaaaaaaa"},
			want: "[]",
		},
		{
			name: "Test aaababa",
			args: args{"aaababa"},
			want: "[1,3]",
		},
		{
			name: "Test caababa",
			args: args{"caababa"},
			want: "null",
		},
		{
			name: "Test rotaratov",
			args: args{"rotaratov"},
			want: "[4,8]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge11.GetIndexsForPalindrome(tt.args.word); got != tt.want {
				t.Errorf("GetIndexsForPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
