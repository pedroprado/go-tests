package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := [][]string{{"0", "a"}, {"0", "b"}, {"2", "c"}, {"1", "d"}}
	countSort(arr)

	// fmt.Println("Result : ", result)
}

func countSort(arr [][]string) {
	// Write your code here
	freqArr := make([][]string, 100)

	half := len(arr) / 2

	for i, v := range arr {
		n, _ := strconv.Atoi(v[0])
		valueConverted := v[1]
		if i < half {
			valueConverted = "-"
		}

		freqArr[n] = append(freqArr[n], valueConverted)
	}

	for _, v := range freqArr {
		if len(v) > 0 {
			for _, s := range v {
				fmt.Printf("%s ", s)
			}

		}
	}

	fmt.Println()
}
