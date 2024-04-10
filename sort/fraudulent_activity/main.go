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

type TrailingDays []int32

func (trailingDays TrailingDays) Len() int {
	return len(trailingDays)
}

func (trailingDays TrailingDays) Less(i, j int) bool {
	return trailingDays[i] < trailingDays[j]
}

func (trailingDays TrailingDays) Swap(i, j int) {
	trailingDays[i], trailingDays[j] = trailingDays[j], trailingDays[i]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	buf := make([]byte, 1000)

	expenditure := []int32{}

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}

			break
		}

		numbersString := string(buf[0:n])
		// numbersString = strings.Trim(numbersString, "\n")
		numbers := strings.Split(numbersString, " ")
		for _, number := range numbers {
			if number != "" && number != " " {
				nu, err := strconv.ParseInt(number, 10, 32)
				if err != nil {
					log.Fatal(err)
				}

				expenditure = append(expenditure, int32(nu))
			}
		}
	}

	fmt.Println("expenditure: ", len(expenditure))
	trailingDays := int32(10000)

	// expenditure := []int32{10, 20, 30, 40, 50}
	// trailingDays := int32(3)

	start := time.Now()
	fmt.Println("Started calculation")
	numberOfNotifications := activityNotifications(expenditure, trailingDays)

	fmt.Println("Ended calculation in seconds: ", time.Since(start).Seconds())
	fmt.Println("Number of notifications: ", numberOfNotifications)
}

// number of notifications sent
func activityNotifications(expenditure []int32, d int32) int32 {
	if d == 0 || len(expenditure) <= int(d) {
		return 0
	}

	var notifications int32

	trailingDays := make([]int32, d)
	for i := 0; i < int(d); i++ {
		trailingDays[i] = expenditure[i]
	}
	sort.Sort(TrailingDays(trailingDays))

	for i := int(d); i < len(expenditure); i++ {
		dayExpenditure := expenditure[i]

		median := getMedian(trailingDays)
		if float64(dayExpenditure) >= 2*median {
			notifications++
		}

		// fmt.Println("Traling days: ", trailingDaysOrdered)
		// fmt.Println("Median: ", median)
		// fmt.Println("Day expenditure: ", dayExpenditure)
		elementToAdd := dayExpenditure
		elementToRemove := expenditure[i-int(d)]

		trailingDays = updateTrailingDays(trailingDays, elementToRemove, elementToAdd)

	}

	return notifications
}

func updateTrailingDays(trailingDays []int32, elementToRemove, elementToAdd int32) []int32 {
	trailingDays = removeElement(trailingDays, elementToRemove)
	// fmt.Println(trailingDays[100:110])

	if elementToAdd <= trailingDays[0] {
		newTrailingDays := append([]int32{elementToAdd}, trailingDays[0:]...)
		return newTrailingDays
	}
	if elementToAdd >= trailingDays[len(trailingDays)-1] {
		trailingDays = append(trailingDays, elementToAdd)
		return trailingDays
	}

	return addElement(trailingDays, elementToAdd)
}

func removeElement(arr []int32, elementToRemove int32) []int32 {
	indexToRemove := findIndex(arr, elementToRemove)

	newArr := append([]int32{}, arr[:indexToRemove]...)
	newArr = append(newArr, arr[indexToRemove+1:]...)
	return newArr
}

func addElement(trailingDays []int32, elementToAdd int32) []int32 {
	indexToAdd := findIndex(trailingDays, elementToAdd)
	newTrailingDays1 := append([]int32{}, trailingDays[0:indexToAdd]...)
	newTrailingDays2 := append([]int32{elementToAdd}, trailingDays[indexToAdd:]...)
	newTrailingDays := append(newTrailingDays1, newTrailingDays2...)

	return newTrailingDays
}

func findIndex(arr []int32, elementToAdd int32) int {
	return int(binarySearch(arr, elementToAdd))
}

func binarySearch(arr []int32, elementToRemove int32) int32 {
	var left = 0
	var right = len(arr) - 1
	var middle = 0

	for left != right {
		middle = (left + right) / 2

		if elementToRemove > arr[middle] {
			left = middle + 1
		} else {
			right = middle
		}
	}

	return int32(left)
}

func getMedian(arr []int32) float64 {
	if len(arr)%2 == 0 {
		mean := len(arr) / 2
		median := (float64(arr[mean-1]) + float64(arr[mean])) / 2

		return median
	}

	mean := len(arr) / 2
	median := float64(arr[mean])
	return median
}
