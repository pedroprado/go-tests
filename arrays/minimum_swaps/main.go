package main

import (
	"fmt"
	"strconv"
)

// Minimium number of swaps = minimum number of non-intersecting cycles

// Graphs solution

func main() {
	arr := []int32{7, 1, 3, 2, 4, 5, 6}
	// change 1 {6,1,3,2,4,5,7}
	// change 2 {5,1,3,2,4,6,7}
	// change 3 {4,1,3,2,5,6,7}
	// change 4 {2,1,3,4,5,6,7}
	// change 5 {1,2,3,4,5,6,7}
	swaps := minimumSwapsHash(arr)

	fmt.Println(arr)
	fmt.Println(swaps)

	f, err := strconv.ParseFloat("3,59", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

}

// brute force efficient solution = if the value in index i is not correct, change arr[i] = arr[value]
// Slow for large arrays
func minimumSwapsBruteForce(arr []int32) int32 {
	var count int32

	for i, _ := range arr {
		done := false

		for !done {
			if arr[i] != int32(i)+1 {
				arr = swap(i, int(arr[i])-1, arr)
				count++
				fmt.Println(arr)
			}
			if arr[i] == int32(i)+1 {
				done = true
			}
		}
	}

	return count
}

// Efficient solution
func minimumSwapsHash(arr []int32) int32 {
	visited := make([]bool, len(arr))
	hash := map[int]int32{}

	var count int32

	for i, v := range arr {
		visited[i] = false
		hash[i] = v
	}

	for index, val := range hash {
		currentVal := val
		currentIndex := index
		// we only true to swap not visited elements: visited were already swapped (if needed)
		if visited[currentIndex] == false {
			visited[currentIndex] = true

			// if the element values != than its index, it should be changed
			if currentVal-1 != int32(currentIndex) {
				correctIndex := currentVal - 1

				// only change if its counterparty has not been visited
				for !visited[correctIndex] {
					visited[correctIndex] = true
					currentVal = hash[int(correctIndex)]
					count++
					arr = swap(currentIndex, int(correctIndex), arr)
					correctIndex = currentVal - 1
				}
			}
		}
	}

	return count
}

type Node struct {
	Vertix int32
} // Graphs solution

// Graphs solution = faster, more complex
func minimumSwapsGraphs(arr []int32) int32 {

	return 0
}

func swap(i, j int, arr []int32) []int32 {
	arr[i], arr[j] = arr[j], arr[i]

	return arr
}
