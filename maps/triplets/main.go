package main

import (
	"fmt"
	"time"
)

func main() {
	arrSize := 10000
	arr := make([]int64, arrSize)
	for i := 1; i < arrSize; i++ {
		arr[i] = int64(i) % 99
	}
	// arr := []int64{1, 2, 6, 2, 3, 6, 9, 18, 3, 9}

	var ratio int64 = 3

	start := time.Now()
	count := countTriplets3(arr, ratio)

	fmt.Println("Number of triplets: ", count)
	fmt.Println("Time to execute: ", time.Since(start))

	// fmt.Println(4 % 4)
	// arr0 := []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Println(len(arr0[5:]))
}

// force brute = n^3 time complexity
func countTriplets(arr []int64, r int64) int64 {
	if len(arr) < 3 {
		return 0
	}

	indexMap := map[int64][]int{}
	for i, v := range arr {
		indexMap[v] = append(indexMap[v], i)
	}

	var count int64

	for i := 0; i <= len(arr)-3; i++ {
		first := arr[i]

		for j := i + 1; j <= len(arr)-2; j++ {
			second := arr[j]

			if second == first*r {
				for k := j + 1; k <= len(arr)-1; k++ {
					third := arr[k]

					if third == second*r {
						count++
					}
				}
			}
		}

	}

	return count
}

// force brute = n^2 time complexity with frequency map
func countTriplets2(arr []int64, r int64) int64 {
	if len(arr) < 3 {
		return 0
	}

	frequencyMapRight := map[int64]int{}
	frequencyMapLeft := map[int64]int{}
	for _, v := range arr {
		frequencyMapRight[v] = frequencyMapRight[v] + 1
	}

	var count int64

	for i := 0; i < len(arr); i++ {
		var (
			middle    int64 = arr[i]
			left      int64
			right     int64 = middle * r
			freqLeft  int
			freqRight int
		)

		//there are no triple possible
		if middle%r == 0 {
			left = middle / r
			freqLeft = frequencyMapLeft[left]
		}

		frequencyMapRight[arr[i]] = frequencyMapRight[arr[i]] - 1

		freqRight = frequencyMapRight[right]

		count = count + int64(freqLeft)*int64(freqRight)
		frequencyMapLeft[arr[i]] = frequencyMapLeft[arr[i]] + 1

	}

	return count
}

// indexes map: n^2 time complexity
func countTriplets3(arr []int64, r int64) int64 {
	if len(arr) < 3 {
		return 0
	}

	var count int64

	indexMap := map[int64][]int{}

	for i, v := range arr {
		indexMap[v] = append(indexMap[v], i)
	}

	//all equals: triples == combination 3 to 3
	if len(indexMap) == 1 {
		return combinationsOf3(int64(len(arr)))
	}

	for i, v := range arr {
		firstIndex := i
		first := v
		second := first * r
		third := second * r

		secondIndexes := indexMap[second]
		// there is second index j that satisfy  i < j < k
		if len(secondIndexes) == 0 {
			continue
		}

		lastSecondIndex := secondIndexes[len(secondIndexes)-1]
		if lastSecondIndex < firstIndex {
			continue
		}

		thirdIndexes := indexMap[third]
		if len(thirdIndexes) == 0 {
			continue
		}

		lastThirdIndex := thirdIndexes[len(thirdIndexes)-1]
		for _, secondIndex := range secondIndexes {

			//there is no third index k that satisfy i < j < k
			if lastThirdIndex < secondIndex {
				break
			}

			if secondIndex > firstIndex {
				for k, thirdIndex := range thirdIndexes {
					// find first index with  k > j
					if thirdIndex > secondIndex {
						number := len(thirdIndexes[k:])
						count = count + int64(number)
						break
					}
				}
			}
		}

	}

	return count
}

func areTriplets(first, second, third int64, ratio int64) bool {
	return first*ratio == second && second*ratio == third
}

// number of combinations of n 3-to-3 without repetion
func combinationsOf3(n int64) int64 {
	var total int64

	for i := n - 1; i >= 1; i-- {
		total = total + combinationsOf2(i)
	}

	return total
}

// number of combinations of n 2-to-2 without repetion
func combinationsOf2(n int64) int64 {
	var total int64

	for i := n - 1; i > 0; i-- {
		total += i
	}

	return total
}
