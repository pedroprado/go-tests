package main

import "fmt"

func main() {
	arr := []int32{1, 2, 2, 1, 9, 2, 7, 3, 6, 8}

	countSwaps(arr)
}

func countSwaps(a []int32) {

}

func bubbleSort(arr []int32) {
	count := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				count++
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", count)
	fmt.Printf("First Element: %d\n", arr[0])
	fmt.Printf("Last Element: %d\n", arr[len(arr)-1])
}
