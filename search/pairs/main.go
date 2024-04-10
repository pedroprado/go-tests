package main

import (
	"fmt"
	"time"
)

func main() {
	k := int32(1)
	arr := []int32{1, 2, 3, 4}

	start := time.Now()
	fmt.Println("Starting process")
	result := pairs(k, arr)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func pairs(k int32, arr []int32) int32 {
	// Write your code here

	indexMap := map[int32][]int{}
	for i, v := range arr {
		indexMap[v] = []int{i}
	}

	var count int32
	for _, v := range arr {
		minor := v - k
		indexes := indexMap[minor]
		if len(indexes) == 0 {
			continue
		}

		count++
	}

	return count
}
