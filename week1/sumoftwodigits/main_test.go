package main

import "testing"

func Test_sumOfTwoDigits(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Valid", args{a: 9, b: 7}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfTwoDigits(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("sumOfTwoDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
