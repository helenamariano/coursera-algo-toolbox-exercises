package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Exercise 2
// Maximum Pairwise Product
// Find the maximum product of two distinct numbers in a sequence of non-negative integers.
// Input: A sequence of non-negative integers.
// Output: The maximum value that can be obtained by multiplying two different elements from the sequence.
// E.g.
// Input: 1 2 3
// Output: 6

func main() {

	args, err := parseArguments()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	start := time.Now()
	result, _ := maxPairWiseProductSlow(args)
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time: %s\n", elapsed)

	start = time.Now()
	result, _ = maxPairWiseProductFast(args)
	elapsed = time.Since(start)
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

func maxPairWiseProductFastest(sequence []int) (int, error) {
	if len(sequence) < 2 {
		err := errors.New("Invalid sequence length")
		return 0, err
	}

	maxInt := 0
	secondMaxInt := 0

	for i := 0; i < len(sequence); i++ {
		val := sequence[i]
		if val > maxInt {
			secondMaxInt = maxInt
			maxInt = val
		} else if val > secondMaxInt && val < maxInt {
			secondMaxInt = val
		}
	}

	return (maxInt * secondMaxInt), nil
}

func maxPairWiseProductFast(sequence []int) (int, error) {

	if len(sequence) < 2 {
		err := errors.New("Invalid sequence length")
		return 0, err
	}

	maxInt := 0

	for _, i := range sequence {
		if i > maxInt {
			maxInt = i
		}
	}

	secondMaxInt := 0

	for _, i := range sequence {
		if i > secondMaxInt && i < maxInt {
			secondMaxInt = i
		}
	}

	return (maxInt * secondMaxInt), nil
}

func maxPairWiseProductSlow(sequence []int) (int, error) {

	maxProduct := 0

	if len(sequence) < 2 {
		err := errors.New("Invalid sequence length")
		return maxProduct, err
	}

	for i := 0; i < len(sequence); i++ {

		a := sequence[i]

		for j := 0; j < len(sequence); j++ {
			b := sequence[j]
			if i != j && a != b {
				product := a * b
				if product > maxProduct {
					maxProduct = product
				}
			}
		}
	}
	return maxProduct, nil
}
