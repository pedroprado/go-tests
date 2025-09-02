package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// READ FROM INPUT
// sorting the right element starting for the left
func main() {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter array size: ")
	arrSize, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err.Error())
	}
	n, err := strconv.Atoi(strings.TrimSpace(string(arrSize)))
	if err != nil {
		log.Fatal(err.Error())
	}

	// fmt.Println("Enter array elements (separate by space): ")
	arrElements, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err.Error())
	}
	arrElementArr := strings.Split(strings.TrimSpace(string(arrElements)), " ")
	arr := []int32{}
	for _, e := range arrElementArr {
		element, err := strconv.Atoi(e)
		if err != nil {
			log.Fatal(err.Error())
		}
		arr = append(arr, int32(element))
	}
	if len(arr) != n {
		log.Fatalf("invalid array of size %d. should be size %d", len(arr), n)
	}

	insertionSort2(int32(n), arr)
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
				break

			} else {
				// shift
				arr[j+1] = arr[j]
				numberOfShifts++
			}

		}
		if k < arr[0] {
			arr[0] = k
			// print(arr)
		}
	}

	print(arr)
	fmt.Println("number of shifts: ", numberOfShifts)
}

func print(arr []int32) {
	str := ""
	for _, e := range arr {
		str = str + strconv.Itoa(int(e)) + " "
	}
	fmt.Println(strings.TrimSpace(str))

}
