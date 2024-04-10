package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	a := []int32{1, 3, 5}
	b := []int32{2, 3}
	c := []int32{1, 2, 3}

	start := time.Now()
	fmt.Println("Starting process")
	result := triplets(a, b, c)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func triplets(a []int32, b []int32, c []int32) int64 {
	sort.Sort(ArrayInt(a))
	sort.Sort(ArrayInt(c))

	var count int64

	mapB := map[int32]int{}

	for _, q := range b {
		_, exists := mapB[q]
		if exists {
			continue
		}

		pCount := findNumberMinorEqualQUnique(a, q)
		rCount := findNumberMinorEqualQUnique(c, q)

		fmt.Println("p count:", pCount)
		fmt.Println("r count:", rCount)

		count += pCount * rCount
		mapB[q]++
	}

	return count
}

type ArrayInt []int32

func (a ArrayInt) Len() int {
	return len(a)
}

func (a ArrayInt) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a ArrayInt) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func findNumberPMinorEqualQUnique(q int32, a []int32) int64 {
	var count int64
	if len(a) == 0 {
		return 0
	}

	for _, p := range a {

		if p <= q {
			count++
		} else {
			break
		}
	}

	return count
}

func findNumberMinorEqualQUnique(arr []int32, x int32) int64 {
	index := binarySearch(arr, x)

	if index == -1 {
		return int64(len(arr))
	}

	if index == 0 {
		return 0
	}

	var count int64 = int64(len(arr[0:index]))

	return count
}

func findNumberRMinorEqualQUnique(q int32, c []int32) int64 {
	var count int64
	if len(c) == 0 {
		return 0
	}

	for _, r := range c {

		if r <= q {
			count++
		} else {
			break
		}
	}

	return count
}

// sorted array
// find the index of first element > x
func binarySearch(arr []int32, x int32) int {
	start := 0
	end := len(arr) - 1

	if arr[len(arr)-1] <= x {
		return -1
	}

	for start != end {
		mid := (start + end) / 2
		eMid := arr[mid]
		if eMid <= x {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}
