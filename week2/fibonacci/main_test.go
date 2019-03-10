package main

import (
	"math/big"
	"testing"
)

func Test_getFibonacciNumberFast(t *testing.T) {

	expected239, _ := new(big.Int).SetString("354224848179261915075", 0)

	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"0", args{n: 0}, big.NewInt(0)},
		{"1", args{n: 1}, big.NewInt(0)},
		{"2", args{n: 2}, big.NewInt(1)},
		{"3", args{n: 3}, big.NewInt(2)},
		{"4", args{n: 4}, big.NewInt(3)},
		{"10", args{n: 10}, big.NewInt(55)},
		{"239", args{n: 100}, expected239},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFibonacciNumberFast(tt.args.n); got.Cmp(tt.want) != 0 {
				t.Errorf("getFibonacciNumberFast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFibonacciNumberSlow(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"0", args{n: 0}, 0},
		{"1", args{n: 1}, 1},
		{"2", args{n: 2}, 1},
		{"3", args{n: 3}, 2},
		{"4", args{n: 4}, 3},
		{"10", args{n: 10}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFibonacciNumberSlow(tt.args.n); got != tt.want {
				t.Errorf("getFibonacciNumberSlow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastDigistOfFibonacciNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{"3", args{n: 3}, 2},
		{"331", args{n: 331}, 9},
		{"327305", args{n: 327305}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastDigistOfFibonacciNumber(tt.args.n); got != tt.want {
				t.Errorf("getLastDigistOfFibonacciNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPisanoPeriod(t *testing.T) {
	type args struct {
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Mod=2", args{m: 2}, 3},
		{"Mod=3", args{m: 3}, 8},
		{"Mod=4", args{m: 4}, 6},
		{"Mod=5", args{m: 5}, 20},
		{"Mod=10", args{m: 10}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPisanoPeriod(tt.args.m); got != tt.want {
				t.Errorf("getPisanoPeriod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getModulusOfFibonacciNumber(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"239 1000", args{n: 239, m: 1000}, 161},
		{"2816213588 239", args{n: 2816213588, m: 239}, 151},
		{"2015 3", args{n: 2015, m: 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getModulusOfFibonacciNumber(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("getModulusOfFibonacciNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastDigitOfSumOfFibonacciNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{n: 0}, 0},
		{"1", args{n: 1}, 1},
		{"2", args{n: 2}, 2},
		{"4", args{n: 4}, 7},
		{"9", args{n: 9}, 8},
		{"100", args{n: 100}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastDigitOfSumOfFibonacciNumber(tt.args.n); got != tt.want {
				t.Errorf("getLastDigitOfSumOfFibonacciNumer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLastDigitOfSumOfFibonacciNumberRange(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 0", args{m: 0, n: 0}, 0},
		{"0 1", args{m: 0, n: 1}, 1},
		{"3 7", args{m: 3, n: 7}, 1},
		{"10 10", args{m: 10, n: 10}, 5},
		{"10 200", args{m: 10, n: 200}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastDigitOfSumOfFibonacciNumberRange(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("getLastDigitOfSumOfFibonacciNumberRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
