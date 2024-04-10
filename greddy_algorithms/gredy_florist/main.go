package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	k := int32(2)
	c := []int32{2, 5, 6}

	start := time.Now()
	fmt.Println("Starting process")
	result := getMinimumCost(k, c)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func getMinimumCost(k int32, c []int32) int32 {
	sort.Sort(FlowersPrice(c))

	if len(c) == int(k) {
		cost := int32(0)
		for _, v := range c {
			cost += v
		}
		return cost
	}

	numberOfBuys := len(c) / int(k)
	extraBuys := len(c) % int(k)

	cost := int32(0)

	for i := 1; i <= int(k); i++ {
		flowIndex := i - 1

		flowerCost := c[flowIndex] * 1
		cost += flowerCost

		for buyNumber := 2; buyNumber <= numberOfBuys; buyNumber++ {
			flowIndex = flowIndex + int(k)
			nextFlowerCost := c[flowIndex] * int32(buyNumber)

			cost += nextFlowerCost
		}
	}

	extraBuysCost := int32(0)
	if extraBuys > 0 {
		for i := 1; i <= extraBuys; i++ {
			flowerIndex := len(c) - i
			flowerCost := c[flowerIndex] * (int32(numberOfBuys) + 1)

			extraBuysCost += flowerCost
		}
	}

	return cost + extraBuysCost
}

type FlowersPrice []int32

func (flowersPrice FlowersPrice) Len() int {
	return len(flowersPrice)
}

// descendent
func (flowersPrice FlowersPrice) Less(i, j int) bool {
	return flowersPrice[i] > flowersPrice[j]
}

func (flowersPrice FlowersPrice) Swap(i, j int) {
	flowersPrice[i], flowersPrice[j] = flowersPrice[j], flowersPrice[i]
}
