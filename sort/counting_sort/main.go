package main

import "fmt"

func main() {

	arr := []int32{12, 4, 1, 12}
	result := countingSort(arr)

	fmt.Println("Result : ", result)
}

func countingSort(arr []int32) []int32 {
	// Write your code here
	freqArr := make([]int32, 100)
	for _, v := range arr {
		freqArr[v]++
	}

	sorted := make([]int32, len(arr))
	k := 0
	for i, v := range freqArr {
		if v != 0 {
			for j := 0; j < int(v); j++ {
				sorted[k] = int32(i)
				k++
			}
		}
	}

	return sorted

}
