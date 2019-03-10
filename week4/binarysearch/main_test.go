package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var benchmarkResult []int

func Example_binarySearch() {
	seq := []int{1, 5, 8, 12, 13}
	toFind := []int{8, 1, 23, 1, 11}

	result := binarySearch(seq, toFind)
	fmt.Printf("%v", result)

	// Output:
	// [2 0 -1 0 -1]
}

func Benchmark_binarySearch(b *testing.B) {

	seq := make([]int, b.N)
	for i := range seq {
		seq[i] = i
	}
	toFind := []int{b.N / 2}

	// assign to package level variable to prevent the compiler
	// from optimising
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		benchmarkResult = binarySearch(seq, toFind)
	}
}

func Benchmark_linearSearch(b *testing.B) {

	seq := make([]int, b.N)
	for i := range seq {
		seq[i] = i
	}
	toFind := []int{b.N / 2}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		benchmarkResult = linearSearch(seq, toFind)
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		seq    []int
		toFind []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{"Size = empty list to search", args{seq: []int{}, toFind: []int{8, 1, 23, 1, 11}}, []int{-1, -1, -1, -1, -1}},
		{"Size = 5", args{seq: []int{1, 5, 8, 12, 13}, toFind: []int{8, 1, 23, 1, 11}}, []int{2, 0, -1, 0, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult := binarySearch(tt.args.seq, tt.args.toFind)
			assert.ElementsMatch(t, gotResult, tt.wantResult)
		})
	}
}
