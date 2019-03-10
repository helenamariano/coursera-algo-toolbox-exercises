package main

import (
	"testing"
)

var benchmarkResult int

func Benchmark_majorityElementRecursive(b *testing.B) {

	seq := []int{2, 3, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 1, 2, 3, 4}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		benchmarkResult = majorityElementRecursive(seq)
	}
}

func Benchmark_majorityElement(b *testing.B) {

	seq := []int{2, 3, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 9, 2, 2, 1, 2, 3, 4}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		benchmarkResult = majorityElement(seq)
	}
}

func Test_majorityElementRecursive(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2 3 9 2 2", args{seq: []int{2, 3, 9, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElementRecursive(tt.args.seq); got != tt.want {
				t.Errorf("isThereAMajorityElementRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_majorityElement(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2 3 9 2 2", args{seq: []int{2, 3, 9, 2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.seq); got != tt.want {
				t.Errorf("isThereAMajorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
