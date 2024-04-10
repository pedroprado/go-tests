package main

import (
	"fmt"
	"math"
)

var (
	arr = [][]int32{
		{0, -4, -6, 0, -7, -6},
		{-1, -2, -6, -8, -3, -1},
		{-8, -4, -2, -8, -8, -6},
		{-3, -1, -2, -5, -7, -4},
		{-3, -5, -3, -6, -6, -6},
		{-3, -6, 0, -8, -6, -7},
	}
)

func main() {

	// arr0 := []int32{1, 2, 3, 4, 5, 6, 7}
	// arr01 := []int32{1, 2, 5, 3, 4, 6, 7, 8}
	// arr02 := []int32{1, 2, 5, 3, 7, 4, 6, 8}
	// arr03 := []int32{1, 2, 5, 3, 7, 8, 4, 6}

	arr2 := []int32{1, 2, 5, 3, 7, 8, 6, 4}

	numBribesFinal(arr2)

	// fmt.Println(1 / 2)

}

func numBribesFinal(q []int32) {
	var numBribes int32
	var maxBribes int32 = 2

	for i := len(q) - 1; i >= 0; i-- {
		actual := q[i]
		shouldBe := (int32(i) + 1)
		if actual-shouldBe > maxBribes {
			fmt.Println("Too chaotic")
			return
		}

		for j := int32(math.Max(0, float64(actual-2))); j < int32(i); j++ {
			if q[j] > q[i] {
				numBribes++
			}
		}
	}

	fmt.Println(numBribes)
}

func numBribes(q []int32) {
	var numBribes int32
	var maxBribes int32 = 2
	for i, v := range q {
		if v > int32(i)+1 {
			bribes := v - (int32(i) + 1)
			if bribes > maxBribes {
				fmt.Println("Too chaotic")
				return
			}

			numBribes = numBribes + bribes
			continue

		}

		subArr := q[i+1:]
		var bribes int32
		for j := range subArr {
			if v > subArr[j] {
				bribes++
				if bribes > maxBribes {
					fmt.Println("Too chaotic")
					return
				}
			}
		}

		numBribes = numBribes + bribes

	}
	fmt.Println(numBribes)
}

func rotateArray(arr []int32) []int32 {
	last := arr[0]

	for i := len(arr) - 1; i >= 0; i-- {
		current := arr[i]
		arr[i] = last
		last = current

	}

	return arr
}

func getMaxHG(arr [][]int32) int32 {
	getHGSum := func(arr [][]int32, i, j int) int32 {
		hourGlass := []int32{
			arr[i][j],
			arr[i-1][j-1], arr[i-1][j], arr[i-1][j+1],
			arr[i+1][j-1], arr[i+1][j], arr[i+1][j+1],
		}

		var sum int32 = 0
		for _, v := range hourGlass {
			sum = sum + v
		}

		return sum
	}

	var max int32 = -100
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			sum := getHGSum(arr, i, j)
			fmt.Printf("sum is %v", sum)

			if sum >= max {
				max = sum
				fmt.Printf("new max is %v", max)
			}
		}
	}

	fmt.Printf("max is %v", max)
	return max
}

func repeatedString(s string, n int64) int64 {
	// Write your code here
	segmentSize := int64(len(s))
	substringSize := n
	segmentTimes := substringSize / segmentSize

	segmentAs := 0
	for _, c := range s {
		letter := string(c)
		if letter == "a" {
			segmentAs++
		}
	}

	totalAs := segmentTimes * int64(segmentAs)

	remainderAs := 0
	remainder := substringSize % segmentSize
	if remainder != 0 {
		remainderSubstring := s[:remainder]
		for _, letter := range remainderSubstring {
			if string(letter) == "a" {
				remainderAs++
			}
		}
	}

	return int64(totalAs) + int64(remainderAs)
}
