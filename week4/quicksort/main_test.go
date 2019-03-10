package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var repeatitiveData = []int{2, 3, 9, 2, 9, 1, 1, 2, 4, 2, 10, 1, 1, 2, 3, 2, 2, 5, 2, 2, 0, 1, 2, 5}

func Benchmark_quickSort(b *testing.B) {
	data := []int{2, 3, 9, 2, 2}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		quickSort(data)
	}
}

func Benchmark_quickSortRepeatedElements(b *testing.B) {

	for i := 0; i < b.N; i++ {
		data := copy(repeatitiveData)
		quickSort(data)
	}
}

func Benchmark_quickSort3RepeatedElements(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := copy(repeatitiveData)
		quickSort3(data)
	}
}

func Test_quickSortStress(t *testing.T) {

	const N = 1000
	const LEN = 10000

	for i := 0; i < N; i++ {
		seq := generateRandomNumber(LEN)
		//t.Logf("Input: %v", seq)
		quickSort(seq)
		//t.Logf("Result: %v", seq)
		checkArrayIsAscending(t, seq)
	}
}

func Test_quickSort3Stress(t *testing.T) {

	const N = 1000
	const LEN = 10000

	for i := 0; i < N; i++ {
		seq := generateRandomNumber(LEN)
		//t.Logf("Input: %v", seq)
		quickSort3(seq)
		//t.Logf("Result: %v", seq)
		checkArrayIsAscending(t, seq)
	}
}

func generateRandomNumber(len int) []int {
	result := make([]int, len)
	for i := 0; i < len; i++ {
		result[i] = rand.Intn(len)
	}

	return result
}

func Test_quickSort(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"2 3 9 2 2", args{data: []int{2, 3, 9, 2, 9}}},
		{"Repeat elements", args{data: []int{2, 3, 9, 2, 9, 1, 1, 2, 4, 2, 10, 1, 1, 2, 3, 2, 2, 5, 2, 2, 0, 1, 2, 5}}},
		{"6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1", args{data: []int{6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1}}},
		{"No elements greater than repeated element (ie. 6)", args{data: []int{6, 2, 6, 1, 5, 6, 6, 3, 4, 6, 6}}},
		{"No elements smaller than repeated element (ie. 6)", args{data: []int{6, 7, 6, 9, 10, 9, 12, 7, 15, 17, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort(tt.args.data)
			checkArrayIsAscending(t, tt.args.data)
		})
	}
}

func Test_quickSort3(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"2 3 9 2 2", args{data: []int{2, 3, 9, 2, 9}}},
		{"Repeat elements", args{data: []int{2, 3, 9, 2, 9, 1, 1, 2, 4, 2, 10, 1, 1, 2, 3, 2, 2, 5, 2, 2, 0, 1, 2, 5}}},
		{"6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1", args{data: []int{6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1}}},
		{"No elements greater than repeated element (ie. 6)", args{data: []int{6, 2, 6, 1, 5, 6, 6, 3, 4, 6, 6}}},
		{"No elements smaller than repeated element (ie. 6)", args{data: []int{6, 7, 6, 9, 10, 9, 12, 7, 15, 17, 6}}},
		{"10 17 13 2 5 8 3 3 2 18 3 7 15 14 11 11 12 17 10 16", args{data: []int{10, 17, 13, 2, 5, 8, 3, 3, 2, 18, 3, 7, 15, 14, 11, 11, 12, 17, 10, 16}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort3(tt.args.data)
			t.Logf("Sorted data: %v", tt.args.data)
			checkArrayIsAscending(t, tt.args.data)
		})
	}
}

func checkArrayIsAscending(t *testing.T, arr []int) {
	for i := 1; i < len(arr); i++ {

		if arr[i] < arr[i-1] {
			t.Logf("Error at element %d\n", i)
			assert.Fail(t, "Array is not ascending")
		}

	}

}
func copy(arr []int) []int {
	result := make([]int, len(arr))
	for i, val := range arr {
		result[i] = val
	}
	return result
}

func Test_quickSort3DijkstraRecurisve(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"2 3 9 2 2", args{data: []int{2, 3, 9, 2, 9}}},
		{"Repeat elements", args{data: []int{2, 3, 9, 2, 9, 1, 1, 2, 4, 2, 10, 1, 1, 2, 3, 2, 2, 5, 2, 2, 0, 1, 2, 5}}},
		{"6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1", args{data: []int{6, 4, 2, 3, 9, 8, 9, 4, 7, 6, 1}}},
		{"No elements greater than repeated element (ie. 6)", args{data: []int{6, 2, 6, 1, 5, 6, 6, 3, 4, 6, 6}}},
		{"No elements smaller than repeated element (ie. 6)", args{data: []int{6, 7, 6, 9, 10, 9, 12, 7, 15, 17, 6}}},
		{"10 17 13 2 5 8 3 3 2 18 3 7 15 14 11 11 12 17 10 16", args{data: []int{10, 17, 13, 2, 5, 8, 3, 3, 2, 18, 3, 7, 15, 14, 11, 11, 12, 17, 10, 16}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			quickSort3Dijkstra(tt.args.data)
			checkArrayIsAscending(t, tt.args.data)
		})
	}
}
