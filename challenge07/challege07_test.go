package challenge07_test

import (
	"adventgo2023/challenge07"
	"testing"
)

func Test_drawGift(t *testing.T) {
	type args struct {
		size   int
		symbol string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Cajas 3D con +",
			args: args{4, "+"},
			want: ` 
	      ####
	     #++##
	    #++#+#
	   ####++#
	   #++#+#
	   #++##
	   ####`,
		},
		{
			name: "Test Cajas 3D con *",
			args: args{5, "*"},
			want: ` 
	       #####
	      #***##
	     #***#*#
	    #***#**#
	   #####***#
	   #***#**#
	   #***#*#
	   #***##
	   #####`,
		},
		{
			name: "Test Cajas 3D con ^",
			args: args{1, "^"},
			want: "#\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := challenge07.DrawGift(tt.args.size, tt.args.symbol); got != tt.want {
				t.Errorf("drawGift() = %v, want %v", got, tt.want)
			}
		})
	}
}
