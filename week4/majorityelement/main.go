package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 2
// Check whether an input sequence contains a majority element
// Input: sequence of elements
// Output: the value of the majority element or 0 if no majority element found

func main() {

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		os.Exit(-1)
	}

	result := majorityElement(args)
	fmt.Printf("result: %v\n", result)

	resultRecursive := majorityElementRecursive(args)
	fmt.Printf("result recursive: %v\n", resultRecursive)

}

func majorityElement(seq []int) int {
	elementCount := make(map[int]int)

	majorityCount := len(seq) / 2
	for _, element := range seq {
		elementCount[element]++
		if elementCount[element] > majorityCount {
			return element
		}
	}
	return 0
}

func majorityElementRecursive(seq []int) int {

	seqLen := len(seq)

	if seqLen == 2 {
		if seq[0] == seq[0] {
			return seq[0]
		}

		return 0

	}
	if seqLen == 1 {
		return seq[0]
	}

	mid := seqLen / 2
	seqLower := seq[0:mid]
	seqUpper := seq[mid:seqLen]
	//fmt.Printf("mid: %d\n", mid)
	//fmt.Printf("lower: %v\n", seqLower)
	//fmt.Printf("upper: %v\n", seqUpper)

	majorityInLower := majorityElementRecursive(seqLower)
	majorityInUpper := majorityElementRecursive(seqUpper)
	//fmt.Printf("majority in lower: %v\n", majorityInLower)
	//fmt.Printf("majority in upper: %v\n", majorityInUpper)

	if majorityInLower == majorityInUpper {
		return majorityInLower
	}

	frequencyLower := getFrequency(seq, majorityInLower)
	frequencyUpper := getFrequency(seq, majorityInUpper)

	//fmt.Printf("frequency in lower: %v\n", frequencyLower)
	//fmt.Printf("frequency in upper: %v\n", frequencyUpper)
	if frequencyLower > mid {
		return majorityInLower
	}

	if frequencyUpper > mid {
		return majorityInUpper
	}

	return 0

}

func getFrequency(seq []int, element int) int {
	count := 0
	for _, val := range seq {
		if element == val {
			count++
		}
	}
	return count
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
