package challenge13_test

import (
	"adventgo2023/challenge13"
	"testing"
)

func Test_calculateTime(t *testing.T) {
	type args struct {
		deliveries []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Returns true -02:20:00",
			args: args{[]string{"00:10:00", "01:00:00", "03:30:00"}},
			want: "-02:20:00",
		},
		{
			name: "Test Returns 00:30:00",
			args: args{[]string{"02:00:00", "05:00:00", "00:30:00"}},
			want: "00:30:00",
		},
		{
			name: "Test Returns -05:29:00",
			args: args{[]string{"00:45:00", "00:45:00", "00:00:30", "00:00:30"}},
			want: "-05:29:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge13.CalculateTime(tt.args.deliveries); got != tt.want {
				t.Errorf("calculateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
