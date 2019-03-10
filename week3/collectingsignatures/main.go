package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Exercise 4
// Collecting Signatures
// Given a set of 𝑛 segments {[𝑎0,𝑏0],[𝑎1,𝑏1],...,[𝑎𝑛−1,𝑏𝑛−1]} with integer coordinates on a line,
// find the minimum number 𝑚 of points such that each segment contains at least one point.
// That is, find a set of integers 𝑋 of the minimum size such that for any segment [𝑎𝑖,𝑏𝑖] there
// is a point 𝑥 ∈ 𝑋 such that𝑎𝑖 ≤𝑥≤𝑏𝑖.

type segment struct {
	a, b int
}
type byB []segment

func (a byB) Len() int           { return len(a) }
func (a byB) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byB) Less(i, j int) bool { return a[i].b < a[j].b }

func main() {

	segments, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Segments:%v\n", segments)

	points := getMinPoints(segments)

	fmt.Printf("Min Number of points: %d\n", len(points))
	fmt.Printf("Points: %v\n", points)
}

func getMinPoints(segments []segment) []int {

	// TODO: Check bounds

	// sort by right coordinate asecending
	sort.Sort(byB(segments))

	// greedy choice:
	// right most coord is the optimal point
	point := segments[0].b
	points := []int{point}

	for i := 1; i < len(segments); i++ {
		s := segments[i]
		if point < s.a {
			point = s.b
			points = append(points, point)
		}
	}
	return points
}

func parseArguments() ([]segment, error) {

	var segments []segment

	args, err := getArgsFromCommandLine(-1)
	if err != nil {
		return segments, err
	}

	// Format of input:
	// n (num of of segments)
	// a1, b1, a2, b2 ... an, bn (coords of segments)
	numSegments := args[0]
	if len(args) != ((numSegments * 2) + 1) {
		return segments, errors.New("Unexpected number of arguments")
	}

	for i := 1; i < len(args); i += 2 {
		s := segment{a: args[i], b: args[i+1]}
		segments = append(segments, s)
	}

	return segments, nil
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
