package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	arr := [][]int32{{5, 1}, {2, 1}, {1, 1}, {8, 1}, {10, 0}, {5, 0}}
	k := int32(3)
	// arr := [][]int32{{13, 1}, {10, 1}, {9, 1}, {8, 1}, {13, 1}, {12, 1}, {18, 1}, {13, 1}}
	// k := int32(5)
	start := time.Now()
	fmt.Println("Starting process")
	result := luckBalance(k, arr)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)
}

// what is the maximum amount of luck for k important contests?
func luckBalance(k int32, contests [][]int32) int32 {
	// Write your code here

	if len(contests) == 0 {
		return 0
	}

	luckSumLosingAllContests := int32(0)
	numberOfImportantContests := int32(0)
	importantContestsLuck := []int32{}
	notImportantContestsLuck := []int32{}
	for i := 0; i < len(contests); i++ {
		luckSumLosingAllContests += contests[i][0]
		numberOfImportantContests += contests[i][1]

		if contests[i][1] == 1 {
			importantContestsLuck = append(importantContestsLuck, contests[i][0])
		} else {
			notImportantContestsLuck = append(notImportantContestsLuck, contests[i][0])
		}
	}

	if numberOfImportantContests <= k {
		return luckSumLosingAllContests
	}

	sort.Sort(Lucks(importantContestsLuck))

	fmt.Println("important contests")
	fmt.Println(importantContestsLuck)
	fmt.Println("not important contests")
	fmt.Println(notImportantContestsLuck)

	sumOfLowestKImportantContests := importantContestsLuck[0]
	for i := 1; i < len(importantContestsLuck)-int(k); i++ {
		newSum := sumOfLowestKImportantContests + importantContestsLuck[i]
		sumOfLowestKImportantContests = newSum
	}

	return luckSumLosingAllContests - 2*sumOfLowestKImportantContests
}

type Lucks []int32

func (ref Lucks) Len() int {
	return len(ref)
}

func (ref Lucks) Less(i, j int) bool {
	return ref[i] < ref[j]
}

func (ref Lucks) Swap(i, j int) {
	ref[i], ref[j] = ref[j], ref[i]
}
