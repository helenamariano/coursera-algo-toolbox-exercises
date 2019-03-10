package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
)

// Exercise 3
// Quicksort with 3-way partition
// Input: sequene to sort
// Output: sequence in decreasing order

func main() {

	data, err := getArgsFromCommandLine(-1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("Data not sorted: %v\n", data)
	quickSort(data)
	fmt.Printf("Data sorted: %v\n", data)

}

func quickSort3(data []int) {
	quickSort3Recurisve(data, 0, len(data)-1)
}

func quickSort3Recurisve(data []int, l, r int) {
	if l >= r {
		return
	}

	m1, m2 := partition3(data, l, r)
	quickSortRecurisve(data, l, m1-1)
	quickSortRecurisve(data, m2+1, r)
}

func quickSort3Dijkstra(data []int) {
	quickSort3Recurisve(data, 0, len(data)-1)
}

func quickSort3DijkstraRecurisve(data []int, l, r int) {
	if l >= r {
		return
	}
	m1, m2 := partitionDijkstra(data, l, r)
	quickSortRecurisve(data, l, m1-1)
	quickSortRecurisve(data, m2+1, r)
}

func partitionDijkstra(data []int, l, r int) (m1, m2 int) {
	pivot := data[l]

	m1 = l
	m2 = r

	for i := l; i < m2; {
		if data[i] < pivot {
			swap(data, i, m1)
			m1++
			i++
		} else if data[i] > pivot {
			m2--
			swap(data, i, m2)
		} else {
			i++
		}

	}

	fmt.Printf("l: %d, r: %d, m1: %d, m2 %d\n", l, r, m1, m2)

	return m1, m2
}

func partition3(data []int, l, r int) (m1, m2 int) {

	pivot := data[l]

	m1 = l
	m2 = l
	for i := l + 1; i <= r; i++ {
		if data[i] < pivot {
			m1++
			swap(data, i, m1)
			if m1 == m2 {
				m2++
				swap(data, i, m2)
			}

		} else if data[i] == pivot {
			if m2 < m1 {
				m2 = m1
			}
			m2++

			swap(data, i, m2)
		}

	}
	swap(data, l, m1)
	return m1, m2
}

func quickSort(data []int) {
	quickSortRecurisve(data, 0, len(data)-1)
}

func quickSortRecurisve(data []int, l, r int) {

	if l >= r {
		return
	}

	m := partition(data, l, r)

	//fmt.Printf("partition: %d\n", m)
	//fmt.Printf("Data: %v\n", data)
	quickSortRecurisve(data, l, m-1)
	quickSortRecurisve(data, m+1, r)
}

func partition(data []int, l, r int) int {
	// take the element at l to be the pivot point

	//fmt.Printf("partition: %d\n", m)
	pivot := data[l]
	j := l

	for i := l + 1; i <= r; i++ {
		if data[i] <= pivot {
			j++
			swap(data, i, j)
		}

	}
	swap(data, l, j)
	return j
}

func swap(data []int, i, j int) {
	tmp := data[i]
	data[i] = data[j]
	data[j] = tmp
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
