package main

import (
	"fmt"
	"strconv"
	"strings"
)

// sorting the right most element
func main() {

	n := int32(14)
	arr := []int32{2, 3, 5, 9, 13, 22, 27, 35, 46, 51, 55, 83, 87, 1}

	insertionSort1(n, arr)

}

func insertionSort1(n int32, arr []int32) {
	// Write your code here

	k := arr[n-1]
	for i := n - 2; i >= 0; i-- {
		if arr[i] < k {
			arr[i+1] = k
			print(arr)
			break

		} else {
			arr[i+1] = arr[i]
			print(arr)
		}

	}
	if k < arr[0] {
		arr[0] = k
		print(arr)
	}
}

func print(arr []int32) {
	str := ""
	for _, e := range arr {
		str = str + strconv.Itoa(int(e)) + " "
	}
	fmt.Println(strings.TrimSpace(str))
}
