package main

import (
	"reflect"
	"testing"
)

func TestGetNumSegmentsContainsPointsNaive(t *testing.T) {
	type args struct {
		segments []int
		points   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"segments = 0 5, 7 10; points = 1, 6, 11", args{segments: []int{0, 5, 7, 10}, points: []int{1, 6, 11}}, []int{1, 0, 0}},
		{"segments = -10, 10; points = -100, 100, 0", args{segments: []int{-10, 10}, points: []int{-100, 100, 0}}, []int{0, 0, 1}},
		{"segments = 0 5, -3 2, 7 10; points =1, 6", args{segments: []int{0, 5, -3, 2, 7, 10}, points: []int{1, 6}}, []int{2, 0}},
		{"segments = 0 5, 7 10; points = 1,1,1", args{segments: []int{0, 5, 7, 10}, points: []int{1, 1, 1}}, []int{1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumSegmentsContainsPointsNaive(tt.args.segments, tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumSegmentsContainsPointsNaive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNumSegmentsContainsPointsFast(t *testing.T) {
	type args struct {
		segments []int
		points   []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"segments = 0 5, 7 10; points = 1, 6, 11", args{segments: []int{0, 5, 7, 10}, points: []int{1, 6, 11}}, []int{1, 0, 0}},
		{"segments = -10, 10; points = -100, 100, 0", args{segments: []int{-10, 10}, points: []int{-100, 100, 0}}, []int{0, 0, 1}},
		{"segments = 0 5, -3 2, 7 10; points =1, 6", args{segments: []int{0, 5, -3, 2, 7, 10}, points: []int{1, 6}}, []int{2, 0}},
		{"segments = 0 5, 7 10; points = 1,1,1", args{segments: []int{0, 5, 7, 10}, points: []int{1, 1, 1}}, []int{1, 1, 1}},
		{"segments = 1 2, 1 2; points = 1, 2", args{segments: []int{1, 2, 1, 2}, points: []int{1, 2}}, []int{2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNumSegmentsContainsPointsFast(tt.args.segments, tt.args.points); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNumSegmentsContainsPointsFast() = %v, want %v", got, tt.want)
			}
		})
	}
}
