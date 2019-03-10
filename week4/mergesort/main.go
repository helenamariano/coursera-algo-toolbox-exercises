package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := getArgsFromCommandLine(-1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Input: %v\n", data)
	sortedData, numInversions := getNumInversions(data)
	fmt.Printf("Result: %v\n", sortedData)
	fmt.Printf("Num Inversions: %v\n", numInversions)
}

func getNumInversions(data []int) ([]int, int) {
	numInversions := 0
	if len(data) == 1 {
		return data, numInversions
	}

	mid := len(data) / 2
	a, numInversionsA := getNumInversions(data[:mid])
	b, numInversionsB := getNumInversions(data[mid:len(data)])

	ab, numInversionsAB := getNumInversionsMerge(a, b)

	fmt.Printf("a: %v, b: %v, ab: %v\n", a, b, ab)
	numInversions += numInversionsA + numInversionsB + numInversionsAB
	fmt.Printf("inva: %v, invb: %v, invab: %v, total: %v\n", numInversionsA, numInversionsB, numInversionsAB, numInversions)
	return ab, numInversions
}

func getNumInversionsMerge(a, b []int) ([]int, int) {
	numInversions := 0

	c := make([]int, len(a)+len(b))

	ai := 0
	bi := 0
	ci := 0
	for ; ai < len(a) && bi < len(b); ci++ {
		if a[ai] <= b[bi] {
			c[ci] = a[ai]
			ai++
		} else {
			numInversions += (len(a) - ai)
			c[ci] = b[bi]
			bi++
		}
	}
	// copy the remaining elements in a
	for ; ai < len(a); ci++ {
		c[ci] = a[ai]
		ai++
	}
	// copy the remaining elements in b
	for ; bi < len(b); ci++ {
		c[ci] = b[bi]
		bi++
	}
	return c, numInversions
}

func mergeSort(data []int) []int {

	if len(data) == 1 {
		return data
	}

	mid := len(data) / 2
	b := mergeSort(data[:mid])
	c := mergeSort(data[mid:len(data)])

	sortedData := merge(b, c)

	return sortedData
}

func merge(a, b []int) []int {

	c := make([]int, len(a)+len(b))

	ai := 0
	bi := 0
	ci := 0
	for ; ai < len(a) && bi < len(b); ci++ {
		if a[ai] <= b[bi] {
			c[ci] = a[ai]
			ai++
		} else {
			c[ci] = b[bi]
			bi++
		}
	}
	// copy the remaining elements in a
	for ; ai < len(a); ci++ {
		c[ci] = a[ai]
		ai++
	}
	// copy the remaining elements in b
	for ; bi < len(b); ci++ {
		c[ci] = b[bi]
		bi++
	}

	return c

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
