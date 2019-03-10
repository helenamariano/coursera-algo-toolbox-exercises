package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 5
// The goal of this problem is to represent a given positive integer n as a sum of as many pairwise distinct
// positive integers as possible. That is, to find the maximum 𝑘 such that 𝑛 can be written as
// 𝑎1+𝑎2+···+𝑎𝑘 where𝑎1,...,𝑎𝑘 arepositiveintegersand𝑎𝑖 ̸=𝑎𝑗 forall1≤𝑖<𝑗≤𝑘.

type segment struct {
	a, b int
}
type byB []segment

func (a byB) Len() int           { return len(a) }
func (a byB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byB) Less(i, j int) bool { return a[i].b < a[j].b }

func main() {

	sum, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Segments:%v\n", sum)

	numValues, err := getValuesForSum(sum)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of integers: %d\n", numValues)
	//fmt.Printf("Integers: %v\n", values)
}

func getValuesForSum(sum int) (int, error) {

	var values []int

	if sum < 1 && sum > 10^9 {
		errMsg := fmt.Sprintf("Invalid value for sum %d", sum)
		return 0, errors.New(errMsg)
	}

	var value int

	for i := 1; sum > 0; i++ {

		if sum <= (2 * i) {
			value = sum
		} else {
			value = i
		}
		//fmt.Printf("Sum: %d\t value: %d\t Values: %v\n", sum, value, values)
		//fmt.Printf("sum %d\t value %d\n", sum, value)
		values = append(values, value)
		sum -= value
	}

	return len(values), nil
}

func parseArguments() (int, error) {

	args, err := getArgsFromCommandLine(1)
	if err != nil {
		return 0, err
	}

	return args[0], nil
}

func getArgsFromCommandLine(expectedNumArgs int) ([]int, error) {
	flag.Parse()
	rawArgs := flag.Args()
	var parsedArgs []int

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
