package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	maxDigits = int32(1000000)
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	buf := make([]byte, 1000000000)
	numbers := []string{}
	arr := []int32{}

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		numbersString := string(buf[:n])
		numbers = strings.Split(numbersString, " ")
		for _, number := range numbers {
			if number != "" && number != " " {
				isNegative := false
				if strings.Contains(number, "-") {
					isNegative = true
					number = string(number[1:])
				}
				nInt, err := strconv.ParseInt(number, 10, 32)
				if err != nil {
					log.Fatal(err)
				}

				if isNegative {
					nInt = -1 * nInt
				}
				arr = append(arr, int32(nInt))
			}
		}

	}

	fmt.Println("12" > "13")
	fmt.Println([]byte("12"), []byte("13"))

	fmt.Println(bytes.Compare([]byte("12"), []byte("13")))

	// fmt.Println("Result: ", result)
}

func bigSorting(unsorted []string) []string {
	// Write your code here

	sorted := []string{}

	lenMap := map[int][]string{}
	for _, v := range unsorted {
		lenMap[len(v)] = append(lenMap[len(v)], v)
	}

	lenList := []int{}
	for k, _ := range lenMap {
		lenList = append(lenList, k)
	}

	insertionSortInt(lenList)

	for _, len := range lenList {
		arrString, _ := lenMap[len]
		insertionSortString(arrString)

		sorted = append(sorted, arrString...)

	}

	return sorted
}

func bubleSortInt(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func bubleSortString(arr []string) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func insertionSortString(arr []string) {
	if len(arr) < 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		k := arr[i]

		for j := i - 1; j >= 0; j-- {
			comparison := bytes.Compare([]byte(arr[j]), []byte(k))

			if comparison == -1 {
				arr[j+1] = k
				break

			} else {
				if comparison == 0 {
					continue
				}
				arr[j+1] = arr[j]

			}

		}
		if k < arr[0] {
			arr[0] = k
		}

	}

}

func insertionSortInt(arr []int) {
	if len(arr) < 1 {
		return
	}

	for i := 1; i < len(arr); i++ {
		k := arr[i]

		for j := i - 1; j >= 0; j-- {
			if arr[j] < k {
				arr[j+1] = k
				break

			} else {
				if arr[j] == k {
					continue
				}
				arr[j+1] = arr[j]

			}

		}
		if k < arr[0] {
			arr[0] = k
		}

	}
}
