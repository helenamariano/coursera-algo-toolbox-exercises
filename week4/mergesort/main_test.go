package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeSortStress(t *testing.T) {

	const N = 1000
	const LEN = 10000

	for i := 0; i < N; i++ {
		seq := generateRandomNumber(LEN)
		result := mergeSort(seq)
		checkIsAscending(t, result)
	}
}

func Test_mergeSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"7 2 5 3 7 13 1 6", args{data: []int{7, 2, 5, 3, 7, 13, 1, 6}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mergeSort(tt.args.data)
			checkIsAscending(t, result)

		})
	}
}

func generateRandomNumber(len int) []int {
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[i] = rand.Intn(len)
	}

	return result
}

func checkIsAscending(t *testing.T, arr []int) {
	for i := 1; i < len(arr); i++ {

		if arr[i] < arr[i-1] {
			t.Logf("Error at element %d\n", i)
			assert.Fail(t, "Array is not ascending")
		}

	}
}

func Test_getNumInversions(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"2 3 9 2 9", args{data: []int{2, 3, 9, 2, 9}}, 2},
		{"2 3 9 2 9 8 4", args{data: []int{2, 3, 9, 2, 9, 8, 4}}, 7},
		{"9 8 7 3 2 1", args{data: []int{9, 8, 7, 3, 2, 1}}, 15},
		{"2, 4, 1, 3, 5", args{data: []int{2, 4, 1, 3, 5}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortedResult, numInverions := getNumInversions(tt.args.data)
			checkIsAscending(t, sortedResult)
			assert.Equal(t, tt.want, numInverions)
		})
	}
}
