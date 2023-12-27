package main

import (
	"reflect"
	"testing"
)

func Test_optimizeIntervals(t *testing.T) {
	type args struct {
		intervals []Interval
	}
	tests := []struct {
		name string
		args args
		want []Interval
	}{
		{
			name: "Test [[2, 8]]",
			args: args{[]Interval{{5, 8}, {2, 7}, {3, 4}}},
			want: []Interval{{2, 8}},
		},
		{
			name: "Test [[1, 6], [8, 10]]",
			args: args{[]Interval{{1, 3}, {8, 10}, {2, 6}}},
			want: []Interval{{1, 6}, {8, 10}},
		},
		{
			name: "Test [[1, 2], [3, 4], [5, 6]]",
			args: args{[]Interval{
				{3, 4},
				{1, 2},
				{5, 6},
			}},
			want: []Interval{{1, 2}, {3, 4}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optimizeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optimizeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
