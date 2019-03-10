package main

import (
	"testing"
	"time"
)

func Test_maxPairWiseProductSlow(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sequence of 2", args{sequence: []int{100000, 90000}}, 9000000000},
		{"Sequence of 3 repeated", args{sequence: []int{5, 5, 4}}, 20},
		{"Sequence of 3", args{sequence: []int{1, 2, 3}}, 6},
		{"Sequence of 10", args{sequence: []int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3}}, 140},
		{"Sequence of 200000", args{sequence: generateAsecendingSequence(200000)}, 39999800000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if got, _ := maxPairWiseProductSlow(tt.args.sequence); got != tt.want {
				t.Errorf("maxPairWiseProductSlow() = %v, want %v", got, tt.want)
			}
			elapsed := time.Since(start)
			t.Log("Time: " + elapsed.String())
		})
	}
}

func Test_maxPairWiseProductFast(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sequence of 2", args{sequence: []int{100000, 90000}}, 9000000000},
		{"Sequence of 3 repeated", args{sequence: []int{5, 5, 4}}, 20},
		{"Sequence of 3", args{sequence: []int{1, 2, 3}}, 6},
		{"Sequence of 10", args{sequence: []int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3}}, 140},
		{"Sequence of 200000", args{sequence: generateAsecendingSequence(200000)}, 39999800000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if got, _ := maxPairWiseProductFast(tt.args.sequence); got != tt.want {
				t.Errorf("maxPairWiseProductFast() = %v, want %v", got, tt.want)
			}
			elapsed := time.Since(start)
			t.Log("Time: " + elapsed.String())
		})
	}
}

func generateAsecendingSequence(end int) []int {
	var result []int

	for i := 1; i <= end; i++ {
		result = append(result, i)
	}

	return result
}

func Test_maxPairWiseProductFastest(t *testing.T) {
	type args struct {
		sequence []int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Sequence of 2", args{sequence: []int{100000, 90000}}, 9000000000, false},
		{"Sequence of 3 repeated", args{sequence: []int{5, 5, 4}}, 20, false},
		{"Sequence of 3", args{sequence: []int{1, 2, 3}}, 6, false},
		{"Sequence of 10", args{sequence: []int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3}}, 140, false},
		{"Sequence of 200000", args{sequence: generateAsecendingSequence(200000)}, 39999800000, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			got, err := maxPairWiseProductFastest(tt.args.sequence)
			elapsed := time.Since(start)
			t.Log("Time: " + elapsed.String())
			if (err != nil) != tt.wantErr {
				t.Errorf("maxPairWiseProductFastest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("maxPairWiseProductFastest() = %v, want %v", got, tt.want)
			}
		})
	}
}
