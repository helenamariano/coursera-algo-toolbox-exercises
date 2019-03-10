package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 5
// Organising a lottery
// Given a set of points on a line and a set of segments on a line.
// The goal is to compute, for each point, the number of segments that contain this point.

func main() {

	segments, points, err := parseArguments()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("segments: %v\n", segments)
	fmt.Printf("points: %v\n", points)

	result := GetNumSegmentsContainsPointsNaive(segments, points)
	fmt.Printf("result (naive): %v\n", result)

	result = GetNumSegmentsContainsPointsFast(segments, points)
	fmt.Printf("result (fast): %v\n", result)
}

func parseArguments() (segments, points []int, err error) {
	args, err := getArgsFromCommandLine(-1)

	if err != nil {
		return segments, points, err
	}

	if len(args) < 2 {
		return segments, points, errors.New("Invalid input")
	}

	numSegments := args[0]
	numPoints := args[1]

	if len(args) != (2 + (2 * numSegments) + numPoints) {
		return segments, points, errors.New("Invalid input")
	}
	pointsIndex := 2 + (2 * numSegments)
	segments = args[2:pointsIndex]
	points = args[pointsIndex:len(args)]

	return segments, points, nil
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
