package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	max = int32(10000)
)

// find median using counting sort
func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	buf := make([]byte, 1000000000)

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
		numbers := strings.Split(numbersString, " ")
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

	result := findMedian(arr)

	fmt.Println("Result: ", result)
}

func findMedian(arr []int32) int32 {
	// Write your code here

	freqArr := make([]int32, 2*max)
	for _, v := range arr {
		indexToPut := max + v
		freqArr[indexToPut]++
	}

	sorted := make([]int32, len(arr))
	k := 0
	for i, freq := range freqArr {
		if freq > 0 {
			v := int32(i) - max
			for j := 0; j < int(freq); j++ {
				sorted[k] = v
				k++
			}
		}
	}

	middle := len(arr) / 2
	if len(arr)%2 == 0 {
		return (sorted[middle-1] + sorted[middle]) / 2
	}

	return sorted[middle]
}
