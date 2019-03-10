package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

// Exercise 2
// Implement an algorithm for the fractional knapsack problem
// Input:
// num items, capacity, value-0, weight-0... value-ith, weight-ith
// Output:
// maximal value of fractions of items that fit into the knapsac
// E.g. 3 50 60 20 100 50 120 30 => 180.0000

type knapSackItem struct {
	v, w int
}

func (a knapSackItem) getVwRatio() float64 { return float64(a.v) / float64(a.w) }

type byVwRatio []knapSackItem

func (a byVwRatio) Len() int           { return len(a) }
func (a byVwRatio) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byVwRatio) Less(i, j int) bool { return a[i].getVwRatio() < a[j].getVwRatio() }

// go run main.go 3 50 60 20 100 50 120 30 => 180.0000
// go run main.go 1 10 500 30 => 166.6667

func main() {

	maxCapacity, items, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Max Capacity: %d\n", maxCapacity)
	fmt.Printf("items: %v\n", items)

	result, err := maxValueOfLoot(maxCapacity, items)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Result: %.4f\n", result)
}

func maxValueOfLoot(maxCapacity int, items []knapSackItem) (float64, error) {
	// TODO: check bounds

	// sorted by v/r descending
	sort.Sort(sort.Reverse(byVwRatio(items)))
	fmt.Printf("items sorted: %v\n", items)

	remainingCapacity := float64(maxCapacity)
	value := float64(0)

	for _, item := range items {

		if remainingCapacity == 0 {
			break
		}

		w := math.Min(float64(item.w), remainingCapacity)
		value += (w * item.getVwRatio())
		remainingCapacity -= w
	}

	return value, nil
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

func parseArguments() (int, []knapSackItem, error) {

	var items []knapSackItem

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return 0, items, err
	}

	numItems := args[0]
	maxCapacity := args[1]
	if len(args) != ((numItems * 2) + 2) {
		return 0, items, errors.New("Unexpected number of arguments")
	}

	for i := 2; i < len(args); i += 2 {
		v := args[i]
		w := args[i+1]
		item := knapSackItem{v: v, w: w}
		items = append(items, item)
	}

	return maxCapacity, items, nil
}
