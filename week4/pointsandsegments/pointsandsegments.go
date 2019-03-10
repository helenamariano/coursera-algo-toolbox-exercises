package main

type valTypes int

const (
	left  valTypes = 0
	point valTypes = 1
	right valTypes = 2
	// point value must be between left and right
)

type entry struct {
	val     int
	valType valTypes
}

// GetNumSegmentsContainsPointsNaive is..
func GetNumSegmentsContainsPointsNaive(segments, points []int) []int {

	var counts []int
	for _, p := range points {
		count := 0
		for i := 0; i < len(segments); i += 2 {
			if p >= segments[i] && p <= segments[i+1] {
				count++
			}
		}
		counts = append(counts, count)
	}
	return counts
}

// GetNumSegmentsContainsPointsFast Retreives number of sements which contains a given list points
func GetNumSegmentsContainsPointsFast(segments, points []int) []int {
	list := createEntriesList(segments, points)
	list = mergeSort(list)

	counts := make(map[int]int)

	lCount := 0
	rCount := 0
	for _, entry := range list {
		if entry.valType == left {
			lCount++
		} else if entry.valType == right {
			rCount++
		} else {
			numSegmentsContainingP := lCount - rCount
			counts[entry.val] = numSegmentsContainingP
		}
	}

	// return the result in the same order as the in points were given
	// as the order could have changed when sorted
	var result []int
	for _, p := range points {
		result = append(result, counts[p])
	}

	return result
}

func mergeSort(list []entry) []entry {
	if len(list) == 1 {
		return list
	}

	m := len(list) / 2
	left := list[:m]
	right := list[m:len(list)]

	a := mergeSort(left)
	b := mergeSort(right)

	return merge(a, b)
}

func merge(a, b []entry) []entry {

	var c []entry

	ai := 0
	bi := 0
	for ai < len(a) && bi < len(b) {

		if a[ai].val == b[bi].val {
			if a[ai].valType <= b[bi].valType {
				c = append(c, a[ai])
				ai++
			} else {
				c = append(c, b[bi])
				bi++
			}
		} else if a[ai].val < b[bi].val {
			c = append(c, a[ai])
			ai++
		} else {
			c = append(c, b[bi])
			bi++
		}
	}

	for ; ai < len(a); ai++ {
		c = append(c, a[ai])
	}

	for ; bi < len(b); bi++ {
		c = append(c, b[bi])
	}

	return c
}

func createEntriesList(segments, points []int) []entry {
	var list []entry
	for i := 0; i < len(segments); i += 2 {
		entryl := entry{val: segments[i], valType: left}
		entryr := entry{val: segments[i+1], valType: right}
		list = append(list, entryl)
		list = append(list, entryr)
	}

	for _, p := range points {

		entry := entry{val: p, valType: point}
		list = append(list, entry)
	}
	return list
}
