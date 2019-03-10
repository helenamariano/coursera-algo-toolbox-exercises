package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 1
// Find the minimum number of coins needed to change the input value (an integer) into coins with
// denominations 1, 5, and 10.
// Input: 28
// Output: 6

func main() {

	arg, err := parseArgument()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	result, err := getMinNumberOfCoins(arg)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %d\n", result)
}

// this must be in descending order
var coinDenominations = []int{10, 5, 1}

// performance is O(1)
// it is determined by how many coin denominations there are.
// Since is it fixed (3) then it is constant
func getMinNumberOfCoins(amount int) (int, error) {

	totalNumCoins := 0

	if amount < 1 || amount > 10000 {
		errMsg := fmt.Sprintf("Amount = %d is not between 1 and 10000", amount)
		return totalNumCoins, errors.New(errMsg)
	}

	for _, coin := range coinDenominations {
		fmt.Printf("Amount Left: %d\n", amount)
		if amount == 0 {
			break
		}

		if coin > amount {
			continue
		}

		numCoins := amount / coin
		totalNumCoins += numCoins
		amount -= (numCoins * coin)
	}

	return totalNumCoins, nil
}

func parseArgument() (int, error) {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		errMsg := fmt.Sprintf("Expected 1 argument got %d: %s", len(args), args)
		err := errors.New(errMsg)
		return 0, err
	}
	arg, err := strconv.Atoi(args[0])
	return arg, err
}
