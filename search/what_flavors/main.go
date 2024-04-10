package main

import (
	"fmt"
	"time"
)

func main() {
	money := int32(8)
	cost := []int32{4, 3, 2, 5, 7}
	// arr := [][]int32{{13, 1}, {10, 1}, {9, 1}, {8, 1}, {13, 1}, {12, 1}, {18, 1}, {13, 1}}
	// k := int32(5)
	start := time.Now()
	fmt.Println("Starting process")
	whatFlavors2(cost, money)

	fmt.Println("Time to complete: ", time.Since(start))
}

func whatFlavors(cost []int32, money int32) {
	// Write your code here
	var flavor1 int
	var flavor2 int
	var found bool

	for i := 0; i < len(cost); i++ {
		if found {
			break
		}
		for j := i + 1; j < len(cost); j++ {
			if found {
				break
			}
			if cost[i]+cost[j] == money {
				flavor1 = i
				flavor2 = j
				found = true
			}
		}
	}

	fmt.Printf("%d %d\n", flavor1+1, flavor2+1)
}

func whatFlavors2(cost []int32, money int32) {
	// Write your code here
	var flavor1 int
	var flavor2 int

	indexMap := map[int32][]int{}
	for i, c := range cost {
		if indexMap[c] == nil {
			indexMap[c] = []int{i}
		} else {
			indexMap[c] = append(indexMap[c], i)
		}
	}

	for i, cost1 := range cost {
		cost2 := money - cost1
		indexesCost2 := indexMap[cost2]
		if len(indexesCost2) == 0 {
			continue
		}
		if cost1 == cost2 && len(indexesCost2) < 2 {
			continue
		}
		fmt.Println("cost1: ", cost1)
		fmt.Println("cost2: ", cost2)
		secondIndex := getSecondIndex(cost1, cost2, indexesCost2)

		flavor1 = i
		flavor2 = secondIndex
		break
	}

	fmt.Printf("%d %d\n", flavor1+1, flavor2+1)
}

func getSecondIndex(cost1, cost2 int32, cost2Indexes []int) int {
	if cost1 != cost2 {
		return cost2Indexes[0]
	}

	return cost2Indexes[1]
}
