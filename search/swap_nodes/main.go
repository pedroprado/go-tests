package main

import (
	"fmt"
	"time"
)

// HARD

func main() {
	// 1
	//2  3
	indexes := [][]int32{{2, 3}, {-1, -1}, {-1, -1}}
	queries := []int32{1, 1}

	// 1
	//2  3
	//3 4

	start := time.Now()
	fmt.Println("Starting process")
	result := swapNodes(indexes, queries)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	// Write your code here

	return nil
}
