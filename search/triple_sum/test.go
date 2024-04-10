package main

import "sort"

func triplets2(a []int32, b []int32, c []int32) int64 {
	sort.Sort(ArrayInt(a))
	sort.Sort(ArrayInt(c))

	var count int64

	mapB := map[int32]int{}

	for _, q := range b {
		_, exists := mapB[q]
		if exists {
			continue
		}

		pCount := findNumberPMinorEqualQUnique(q, a)
		rCount := findNumberRMinorEqualQUnique(q, c)
		count += pCount * rCount
		mapB[q]++
	}

	return count
}
