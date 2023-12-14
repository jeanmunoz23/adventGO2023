package challenge14

import "testing"

func TestMaxGifts(t *testing.T) {
	type args struct {
		houses []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Returns 4",
			args: args{[]int{2, 4, 2}},
			want: 4,
		},
		{
			name: "Test Returns 10 (5 + 5)",
			args: args{[]int{5, 1, 1, 5}},
			want: 10,
		},
		{
			name: "Test Returns 9 (4 + 4 + 1)",
			args: args{[]int{4, 1, 1, 4, 2, 1}},
			want: 9,
		},
		{
			name: "Test Returns 103 (3 + 100)",
			args: args{[]int{1, 3, 1, 3, 100}},
			want: 103,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxGifts(tt.args.houses); got != tt.want {
				t.Errorf("MaxGifts() = %v, want %v", got, tt.want)
			}
		})
	}
}
