package main

import (
	"fmt"
	"time"
)

func main() {
	arr := []int32{1, -3, 71, 68, 17}
	start := time.Now()
	fmt.Println("Starting process")
	result := minimumAbsoluteDifference(arr)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)
}

func minimumAbsoluteDifference(arr []int32) int32 {
	// Write your code here

	min := mod(arr[0] - arr[len(arr)-1])
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			absoluteDifference := mod(arr[i] - arr[j])
			if absoluteDifference < min {
				min = absoluteDifference
			}

		}
	}

	return min

}

func mod(i int32) int32 {
	if i < 0 {
		return -1 * i
	}

	return i
}

type Arr []int32

func (arr Arr) Len() int {
	return len(arr)
}

func (arr Arr) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Arr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
