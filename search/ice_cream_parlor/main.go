package main

import (
	"fmt"
)

type Input struct {
	Money int32
	Cost  []int32
}

// solved using linear search
func main() {
	money := int32(100)
	cost := []int32{5, 75, 25}

	fmt.Println("Starting process")
	result := icecreamParlor(money, cost)

	fmt.Println("Result: ", result)

}

func icecreamParlor(m int32, arr []int32) []int32 {
	// Write your code here

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == m {
				return []int32{int32(i + 1), int32(j + 1)}

			}
		}
	}

	return nil
}

func bigSort(arr []int32) []int32 {
	sorted := make([]int32, len(arr))
	freqArr := make([]int32, 10000)

	for _, n := range arr {
		freqArr[n]++
	}

	j := 0
	for indexValue, freq := range freqArr {
		if freq > 0 {
			for i := 0; i < int(freq); i++ {
				sorted[j] = int32(indexValue)
				j++
			}
		}
	}

	return sorted
}
