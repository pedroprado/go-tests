package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	machines := []int64{2, 3}
	goal := int64(5)

	start := time.Now()
	fmt.Println("Starting process")
	result := minTime4(machines, goal)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func minTime(machines []int64, goal int64) int64 {
	// productionByDay := map[int64]int{}

	// for _, machineDays := range machines {
	// 	productionByDay[machineDays]++
	// }

	var productionCount int64
	day := int64(1)

	for productionCount < goal {
		for _, machineDay := range machines {
			if day%machineDay == 0 {
				productionCount++
				if productionCount == goal {
					break
				}
			}
		}

		if productionCount == goal {
			break
		}

		day++
	}

	return day
}

func minTime2(machines []int64, goal int64) int64 {
	productionByDay := map[int64]int64{}

	for _, machineDays := range machines {
		productionByDay[machineDays]++
	}

	var productionCount int64
	day := int64(1)

	for productionCount < goal {
		for machineDay, freq := range productionByDay {
			if day%machineDay == 0 {
				productionCount += freq
				if productionCount >= goal {
					break
				}
			}
		}

		if productionCount >= goal {
			break
		}

		day++
	}

	return day
}

func minTime3(machines []int64, goal int64) int64 {
	// sort.Sort(MachineDays(machines))

	machineProductionInDays := map[int64]int64{}
	for _, machineDays := range machines {
		machineProductionInDays[machineDays]++
	}

	var productionCount int64
	day := int64(1)

	for productionCount < goal {
		multiples := getMultiples(day)

		for _, m := range multiples {
			productionDay, _ := machineProductionInDays[m]
			productionCount += productionDay

			if productionCount >= goal {
				break
			}
		}

		if productionCount >= goal {
			break
		}

		day++
	}

	// fmt.Println(productionByDay)
	return day
}

func minTime4(machines []int64, goal int64) int64 {
	sort.Sort(MachineDays(machines))

	numMachines := int64(len(machines))
	productionDayOfFastestsMachine := machines[0]

	production := int64(0)
	day := int64(1)
	for production < goal {
		if day%productionDayOfFastestsMachine == 0 {
			production += numMachines
		}
		if production >= goal {
			break
		}
		day++
	}

	daysToProduceWithFastestMachine := day

	productionDayOfSlowestMachine := machines[len(machines)-1]
	production = int64(0)
	day = int64(1)
	for production < goal {
		if day%productionDayOfSlowestMachine == 0 {
			production += numMachines
		}
		if production >= goal {
			break
		}
		day++
	}

	daysToProduceWithSlowestMachine := day

	fmt.Println(daysToProduceWithFastestMachine)
	fmt.Println(daysToProduceWithSlowestMachine)

	return day
}

func getMultiples(num int64) []int64 {
	multiples := []int64{}

	if num == 0 {
		return multiples
	}

	for divisor := int64(1); divisor <= num; divisor++ {
		if num%divisor == 0 {
			multiples = append(multiples, divisor)
		}
	}

	return multiples
}

type MachineDays []int64

func (m MachineDays) Len() int {
	return len(m)
}

func (m MachineDays) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MachineDays) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func binarySearch(machines []int64, goal int64) int {
	start := 0
	end := len(machines)

	for start != end {
		mid := (start + end) / 2

		if machines[mid] > goal {
			end = mid
		} else if machines[mid] < goal {
			start = mid + 1
		} else {
			break
		}
	}

	return start
}
