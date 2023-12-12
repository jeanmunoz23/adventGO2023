package challenge08_test

import (
	"adventgo2023/challenge08"
	"testing"
)

func Test_organizeGifts(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "Test organizeGifts('76a11b')",
			args: "76a11b",
			want: "[a]{a}{a}(aaaaaa){b}(b)",
		}, {
			name: "Test organizeGifts('20a')",
			args: "20a",
			want: "{a}{a}",
		}, {
			name: "Test organizeGifts('70b120a4c')",
			args: "70b120a4c",
			want: "[b]{b}{b}[a][a]{a}{a}(cccc)",
		}, {
			name: "Test organizeGifts('9c')",
			args: "9c",
			want: "(ccccccccc)",
		}, {
			name: "Test organizeGifts('19d51e')",
			args: "19d51e",
			want: "{d}(ddddddddd)[e](e)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge08.OrganizeGifts(tt.args); got != tt.want {
				t.Errorf("organizeGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}
