package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	// k := int32(5)
	// arr := []int32{4504, 1520, 5857, 4094, 4157, 3902, 822, 6643, 2422, 7288, 8245, 9948, 2822, 1784, 7820, 3142, 9739, 5629, 5413, 7232}

	// start := time.Now()
	// fmt.Println("Starting process")
	// result := maxMin(k, arr)

	// fmt.Println("Time to complete: ", time.Since(start))
	// fmt.Println("Result: ", result)

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	k := int32(kTemp)

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	start := time.Now()
	result := maxMin2(k, arr)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

// solution 1: n^4 (very slow)
func maxMin(k int32, arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	if len(arr) == 2 {
		return absolute(arr[0] - arr[1])
	}

	minDiff := int64(10000000000)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			left := arr[i]
			right := arr[j]

			arr2 := getArr2(left, right, k, arr)

			diff := arr2[len(arr2)-1] - arr2[0]
			if int64(diff) < minDiff {
				minDiff = int64(diff)
			}
		}
	}

	return int32(minDiff)
}

func absolute(a int32) int32 {
	if a < 0 {
		return a * -1
	}
	return a
}

func getArr2(a, b, k int32, arr []int32) []int32 {
	arr2 := []int32{}
	if a <= b {
		arr2 = append(arr2, a)
		arr2 = append(arr2, b)
	} else {
		arr2 = append(arr2, b)
		arr2 = append(arr2, a)
	}

	arrUpdated := removeElement(a, arr)
	arrUpdated = removeElement(b, arrUpdated)
	for i := 1; i <= int(k)-2; i++ {
		left := arr2[0]
		right := arr2[len(arr2)-1]

		elemBetween, exists := elementBetween(left, right, arrUpdated)
		if exists {
			arr2 = addMiddle(elemBetween, arr2)
			arrUpdated = removeElement(elemBetween, arrUpdated)
			continue
		}

		leftElem, leftDiff := findMinDiffLeft(left, arrUpdated)
		rightElem, rightDiff := findMinDiffRight(right, arrUpdated)

		if leftDiff <= rightDiff {
			arr2 = addLeft(leftElem, arr2)
			arrUpdated = removeElement(leftElem, arrUpdated)
		} else {
			arr2 = addRight(rightElem, arr2)
			arrUpdated = removeElement(rightElem, arrUpdated)
		}

	}

	fmt.Println(arr2)

	return arr2
}

func elementBetween(left, right int32, arr []int32) (int32, bool) {
	for _, v := range arr {
		if v >= left && v <= right {
			return v, true
		}
	}

	return 0, false
}

func removeElement(elem int32, arr []int32) []int32 {
	newArr := []int32{}
	removed := false
	for _, v := range arr {
		if v == elem && !removed {
			removed = true
			continue
		}
		newArr = append(newArr, v)
	}

	return newArr
}

func findMinDiffLeft(a int32, arr []int32) (int32, int32) {
	elem := arr[0]
	minDiff := int64(10000000000)

	for i := 0; i < len(arr); i++ {
		if arr[i] > a {
			continue
		}
		diff := a - arr[i]
		if int64(diff) < minDiff {
			elem = arr[i]
			minDiff = int64(diff)
		}
	}

	return elem, int32(minDiff)
}

func findMinDiffRight(a int32, arr []int32) (int32, int32) {
	elem := arr[0]
	minDiff := int64(10000000000)

	for i := 0; i < len(arr); i++ {
		if arr[i] < a {
			continue
		}
		diff := arr[i] - a
		if int64(diff) < minDiff {
			elem = arr[i]
			minDiff = int64(diff)
		}
	}

	return elem, int32(minDiff)
}

func addMiddle(a int32, arr []int32) []int32 {
	newArr := []int32{arr[0]}
	newArr = append(newArr, a)
	newArr = append(newArr, arr[1:]...)

	return newArr
}

func addLeft(a int32, arr []int32) []int32 {
	newArr := []int32{a}
	newArr = append(newArr, arr...)

	return newArr

}

func addRight(a int32, arr []int32) []int32 {
	arr = append(arr, a)
	return arr
}

// solution 2: greedy
func maxMin2(k int32, arr []int32) int32 {
	sort.Sort(ArrInt(arr))

	min := int64(1000000000)

	for i := 0; i < len(arr)-int(k-1); i++ {
		first := arr[i]
		last := arr[i+int(k-1)]
		diff := last - first
		if int64(diff) < min {
			min = int64(diff)
		}
	}

	return int32(min)
}

type ArrInt []int32

func (arr ArrInt) Len() int {
	return len(arr)
}

func (arr ArrInt) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr ArrInt) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
