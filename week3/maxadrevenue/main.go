package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Exercise 3
// Given two sequences 𝑎1,𝑎2,...,𝑎𝑛 (𝑎𝑖 is the profit per click of the 𝑖-th ad) and
// 𝑏1,𝑏2,...,𝑏𝑛 (𝑏𝑖 is the average number of clicks per day of the 𝑖-th slot), we need
// to partition them into 𝑛 pairs (𝑎𝑖,𝑏𝑗) such that the sum of their products is maximized.
//
// Input: integer n, a-1, a-2.. a-n, b-1, b-2...b-n
// Output Max product (revenue)
// E.g. 1 23 39 => 897

func main() {

	profitPerClicks, clicksPerDays, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Profits Per Clicks %v\n", profitPerClicks)
	fmt.Printf("Clicks Per Days %v\n", clicksPerDays)

	result, err := maxAdRevenue(profitPerClicks, clicksPerDays)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", result)
}

func maxAdRevenue(profitPerClicks, clicksPerDays []int) (int64, error) {

	// Check bounds
	if len(profitPerClicks) != len(clicksPerDays) ||
		len(profitPerClicks) < 1 ||
		len(profitPerClicks) > 10000 {
		return 0, errors.New("Invalid input")
	}

	// sequences descending
	sort.Sort(sort.Reverse(sort.IntSlice(profitPerClicks)))
	sort.Sort(sort.Reverse(sort.IntSlice(clicksPerDays)))

	revenue := int64(0)

	for i := range profitPerClicks {
		revenue += (int64(profitPerClicks[i]) * int64(clicksPerDays[i]))
	}

	return revenue, nil
}

func parseArguments() ([]int, []int, error) {

	var profitPerClicks, clicksPerDays []int

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return profitPerClicks, clicksPerDays, err
	}

	fmt.Printf("Args: %v", args)

	// Format of input:
	// n (num of values in sequence)
	// a1...an (profit per click)
	// b1...bn (num clicks per day)
	numItems := args[0]
	if len(args) != ((numItems * 2) + 1) {
		return profitPerClicks, clicksPerDays, errors.New("Unexpected number of arguments")
	}

	for i := 1; i < len(args); i += 2 {
		profitPerClicks = append(profitPerClicks, args[i])
		clicksPerDays = append(clicksPerDays, args[i+1])
	}

	return profitPerClicks, clicksPerDays, nil
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
