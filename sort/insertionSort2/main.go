package main

import (
	"fmt"
	"strconv"
	"strings"
)

// sorting the right element starting for the left
func main() {

	n := int32(5)
	arr := []int32{2, 1, 3, 1, 2}

	insertionSort2(n, arr)

}

func insertionSort2(n int32, arr []int32) {
	numberOfShifts := int32(0)

	if n == 1 {
		print(arr)
		return
	}

	for i := 1; i < len(arr); i++ {
		k := arr[i]

		for j := i - 1; j >= 0; j-- {
			if arr[j] < k {
				arr[j+1] = k
				// print(arr)
				break

			} else {
				if arr[j] == k {
					continue
				}
				arr[j+1] = arr[j]
				numberOfShifts++

			}

		}
		if k < arr[0] {
			arr[0] = k
			// print(arr)
		}

		fmt.Printf("number of shifts on iteration %d: %d\n", i, numberOfShifts)
		print(arr)

	}

	print(arr)
}

func print(arr []int32) {
	str := ""
	for _, e := range arr {
		str = str + strconv.Itoa(int(e)) + " "
	}
	fmt.Println(strings.TrimSpace(str))
}
