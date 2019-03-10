package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"
)

// Exercise 1
// Fibonacci Number
// Given an integer 𝑛, find the 𝑛th Fibonacci number 𝐹𝑛.
// E.g.
// Input: 10
// Output

// Exercise 2
// Last Digit of a Large Fibonacci Number
// Given an integer 𝑛, find the last digit of the 𝑛th Fibonacci number 𝐹𝑛 (that is, 𝐹𝑛 mod 10).
// E.g.
// Input: 331
// Output: 9

func main() {

	args, err := parseArguments()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	start := time.Now()
	result := getLastDigitOfSumOfFibonacciNumberRange(args[0], args[1])
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}

func parseArguments() ([]int, error) {
	flag.Parse()
	rawArgs := flag.Args()

	var parsedArgs []int

	for i := range rawArgs {
		intValue, err := strconv.Atoi(rawArgs[i])
		if err != nil {
			return parsedArgs, err
		}
		parsedArgs = append(parsedArgs, intValue)
	}
	return parsedArgs, nil
}

func getLastDigitOfSumOfFibonacciNumberRange(m, n int) int {

	var fn, fm int
	a := 0
	b := 1

	if m < 2 {
		fm = m
	} else {
		fm = a + b
	}

	if n < 2 {
		fn = n
	} else {
		fn = a + b
	}

	if n >= 2 {

		nLimit := n + 2
		mLimit := m + 1

		for i := 2; i <= nLimit; i++ {
			fn = (a + b) % 10
			//fn = a + b
			if i == mLimit {
				fm = fn
			}

			//fmt.Printf("i: %d\tfm: %d \tfn: %d\n", i, fm, fn)
			a = b
			b = fn
		}

		if m == n {
			fm = -1
		}
	}

	fmt.Printf("fm: %d \tfn: %d\n", fm, fn)

	result := fn - fm
	if result < 0 {
		result += 10
	}

	return result
}

func getLastDigitOfSumOfFibonacciNumber(n int) int {

	if n < 2 {
		return n
	}

	a := 0
	b := 1
	c := a + b

	for i := 2; i <= n+2; i++ {
		c = (a + b) % 10
		//c = a + b
		a = b
		b = c
	}

	return c - 1
}

func getModulusOfFibonacciNumber(n int, m int) int {

	period := getPisanoPeriod(m)

	rem := n % period
	f := getFibonacciNumberFast(rem)

	fmt.Printf("Period: %d\t F: %v\n", period, f)

	var result = new(big.Int)
	var bigM = big.NewInt((int64(m)))
	result.Mod(f, bigM)
	// assuming here that the result will fit into int
	return int(result.Int64())
}

func getPisanoPeriod(m int) int {
	// According to https://en.wikipedia.org/wiki/Pisano_period
	// the max pisano period <= 6m where m is the modulus

	var a, b, c uint
	a = 0
	b = 1
	c = a + b

	maxPeriod := 6 * m
	for i := 0; i < maxPeriod; i++ {
		c = (a + b) % uint(m)
		a = b
		b = c
		//fmt.Printf("i: %d\ta: %d\t b: %d\n", i, a, b)
		if a == 0 && b == 1 {
			return i + 1
		}
	}
	return 0
}

func getLastDigistOfFibonacciNumber(n int) uint {
	var sequence []uint

	if n == 0 {
		return 0
	}

	sequence = append(sequence, 0)
	sequence = append(sequence, 1)

	for i := 2; i <= n; i++ {
		lastDigit := (sequence[i-1] + sequence[i-2]) % 10
		sequence = append(sequence, lastDigit)
	}

	return sequence[n]
}
func getFibonacciNumberFast(n int) *big.Int {

	a := big.NewInt(0)
	b := big.NewInt(1)
	var c = new(big.Int)

	if n < 2 {
		return a
	}

	for i := 2; i <= n; i++ {
		c.Add(a, b)
		a.Add(a, b)
		a, b = b, a
	}

	return c
}

func getFibonacciNumberSlow(n int) uint {
	if n < 2 {
		return uint(n)
	}

	result := getFibonacciNumberSlow(n-1) + getFibonacciNumberSlow(n-2)
	return result
}
