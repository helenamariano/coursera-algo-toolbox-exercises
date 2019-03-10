package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 1
// Sum of Two Digits Problem
// Compute the sum of two single digit numbers (simple intro)
// Input: Two single digit numbers. Output: The sum of these num- bers.
// Output format. The sum of a and b.

// E.g.
// Input:
// 9 7
// Output:
// 16
func main() {

	args, err := parseArguments()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result := sumOfTwoDigits(args[0], args[1])
	fmt.Printf("Result: %d\n", result)
}

func parseArguments() ([2]int, error) {
	flag.Parse()
	rawArgs := flag.Args()
	var parsedArgs [2]int

	if len(rawArgs) != 2 {
		err := errors.New("Incorrect number of arguments")
		return parsedArgs, err
	}

	for i := range rawArgs {
		intValue, err := strconv.Atoi(rawArgs[i])
		if err != nil {
			return parsedArgs, err
		}
		parsedArgs[i] = intValue
	}
	return parsedArgs, nil
}

func sumOfTwoDigits(a int, b int) int {
	return a + b
}
