package challenge09_test

import (
	"adventgo2023/challenge09"
	"testing"
)

func Test_adjustLights(t *testing.T) {
	type args struct {
		lights []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test cambias la cuarta luz a 🔴",
			args: args{[]string{"🟢", "🔴", "🟢", "🟢", "🟢"}},
			want: 1,
		},
		{
			name: "Test cambias la segunda luz a 🟢 y la tercera a 🔴",
			args: args{[]string{"🔴", "🔴", "🟢", "🟢", "🔴"}},
			want: 2,
		},
		{
			name: "Test ya están alternadas",
			args: args{[]string{"🟢", "🔴", "🟢", "🔴", "🟢"}},
			want: 0,
		},
		{
			name: "Test cambias la segunda luz a 🟢",
			args: args{[]string{"🔴", "🔴", "🔴"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge09.AdjustLights(tt.args.lights); got != tt.want {
				t.Errorf("adjustLights() = %v, want %v", got, tt.want)
			}
		})
	}
}
