package main

import "testing"

func Test_lcm(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"6 8", args{a: 6, b: 8}, 24},
		{"28851538 1183019", args{a: 28851538, b: 1183019}, 1933053046},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lcm(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}
