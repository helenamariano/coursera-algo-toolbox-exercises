package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Exercise 1
// Implement binary search algorithm

// Input
// num in Sequence
// sequence
// num to Find
// values to find

// E.g.
// input:
//  5 1 5 8 12 13 5 8 1 23 1 11
// expected:
// 2 0 -1 0 -1

func main() {

	seq, toFind, err := parseArguments()

	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(-1)
	}

	fmt.Printf("Seq: %v\n", seq)
	fmt.Printf("To Find: %v\n", toFind)

	start := time.Now()
	result := linearSearch(seq, toFind)
	elapsed := time.Since(start)
	fmt.Println("Linear Search")
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Time: %s\n", elapsed)

	start = time.Now()
	result = binarySearch(seq, toFind)
	elapsed = time.Since(start)
	fmt.Println("Binary Search")
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Time: %s\n", elapsed)

}

func binarySearch(seq, toFind []int) (result []int) {

	result = make([]int, len(toFind))
	for i, valToFind := range toFind {

		result[i] = -1 // by default set to not found

		low, high := 0, len(seq)-1

		for high > low {
			mid := low + (high-low)/2

			val := seq[mid]
			if val == valToFind {
				result[i] = mid
				break
			}

			if valToFind > val {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return result
}

func linearSearch(seq, toFind []int) (result []int) {

	result = make([]int, len(toFind))
	for i, valToFind := range toFind {

		result[i] = -1 // by default set to not found

		for j, val := range seq {
			if val == valToFind {
				result[i] = j
				break
			}
		}
	}

	return result
}

func parseArguments() (seq, toFind []int, err error) {

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return seq, toFind, err
	}

	fmt.Printf("Args: %v\n", args)

	seqLen := args[0]
	seq = args[1 : seqLen+1]
	toFind = args[seqLen+2 : len(args)]

	return seq, toFind, nil
}

func getArgsFromCommandLine(expectedNumArgs int) (parsedArgs []int, err error) {
	flag.Parse()
	rawArgs := flag.Args()

	if expectedNumArgs > 1 && len(rawArgs) != expectedNumArgs {
		errMsg := fmt.Sprintf("Expected %d argument(s) got %d: \n%s", expectedNumArgs, len(rawArgs), rawArgs)
		err := errors.New(errMsg)
		return parsedArgs, err
	}

	for _, arg := range rawArgs {
		val, err := strconv.Atoi(arg)
		if err != nil {
			return parsedArgs, err
		}
		parsedArgs = append(parsedArgs, val)

	}
	return parsedArgs, nil
}
