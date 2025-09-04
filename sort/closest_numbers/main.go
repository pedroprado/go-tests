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
	max = int32(10000000)
)

// for findin closes numbers use counting sort
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

	result := closestNumbers(arr)

	fmt.Println("Result: ", result)
}

func closestNumbers(arr []int32) []int32 {
	// Write your code here
	freqArr := make([]int32, 2*max)

	for _, v := range arr {
		if v == 0 {
			freqArr[max] = 0
			continue
		}
		if v < 0 {
			freqArr[max+v]++
			continue
		}
		freqArr[max+v]++
	}

	sorted := make([]int32, len(arr))
	k := 0
	for i, freq := range freqArr {
		if freq > 0 {
			for j := 0; j < int(freq); j++ {
				sorted[k] = getValue(int32(i))
				k++
			}
		}
	}

	fmt.Println(sorted)
	fmt.Println(len(sorted))

	freqMap := map[int32][]int32{}
	for i := 0; i < len(sorted)-1; i++ {
		absDifference := absDifference(sorted[i], sorted[i+1])
		freqMap[absDifference] = append(freqMap[absDifference], sorted[i])
		freqMap[absDifference] = append(freqMap[absDifference], sorted[i+1])

	}

	min := 2 * max
	for absDiff, _ := range freqMap {
		if absDiff < min {
			min = absDiff
		}
	}

	return freqMap[min]

}

func abs(n int32) int32 {
	if n < 0 {
		return -1 * n
	}
	return n
}

func absDifference(n1, n2 int32) int32 {
	if n1 < 0 && n2 < 0 || n1 > 0 && n2 > 0 {
		return abs(abs(n1) - abs(n2))
	}

	return abs(abs(n1) + abs(n2))
}

func getValue(index int32) int32 {
	if index == max {
		return 0
	}
	if index < max {
		return index - max
	}

	return index - max
}
