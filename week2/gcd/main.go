package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Exercise 3
// Greatest Common Divisor
// Given two integers 𝑎 and 𝑏, find their greatest common divisor.
// E.g.
// Input: 28851538 1183019
// Output: 17657

func main() {
	args, err := parseArguments()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	start := time.Now()
	result := gcd(args[0], args[1])
	elapsed := time.Since(start)
	fmt.Printf("Result: %d\n", result)
	fmt.Printf("Time: %s\n", elapsed)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
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
