package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 6
// Compose the largest number out of a set of integers.
// Input: Set of integers
// E.g. 21 2 => 221

type segment struct {
	a, b int
}
type byB []segment

func (a byB) Len() int           { return len(a) }
func (a byB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byB) Less(i, j int) bool { return a[i].b < a[j].b }

func main() {

	digits, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Digits:%v\n", digits)

	num, err := getLargestNumber(digits)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Largest Number: %s\n", num)
	//fmt.Printf("Integers: %v\n", values)
}

func getLargestNumber(digits map[int]string) (string, error) {

	var result string

	for len(digits) > 0 {

		// get largest digit
		maxDigit := getMaxDigit(digits)
		result += digits[maxDigit]
		delete(digits, maxDigit)
	}

	return result, nil
}

func getMaxDigit(digits map[int]string) int {

	maxDigit := 0
	for digit := range digits {
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}

func parseArguments() (map[int]string, error) {

	flag.Parse()

	parsedArgs := make(map[int]string)
	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return parsedArgs, err
	}

	for _, arg := range args {

		val, err := strconv.Atoi(arg)
		//fmt.Printf("arg: %s\n", arg)
		//fmt.Printf("val: %d\n", val)
		if err != nil {
			return parsedArgs, err
		}
		parsedArgs[val] = arg
	}
	return parsedArgs, nil
}

func getArgsFromCommandLine(expectedNumArgs int) ([]string, error) {
	flag.Parse()
	args := flag.Args()

	if expectedNumArgs > 1 && len(args) != expectedNumArgs {
		errMsg := fmt.Sprintf("Expected %d argument(s) got %d: \n%s", expectedNumArgs, len(args), args)
		err := errors.New(errMsg)
		return args, err
	}

	return args, nil
}
