package main

import "fmt"

func main() {

	num := alternate("ab")

	fmt.Println("Result: ", num)

}

func alternate(s string) int32 {
	// Write your code here
	var (
		maxLen       = int32(0)
		charSet      = map[string]int{}
		uniqueString = ""
		temp         = map[string]int{}
	)

	for _, c := range s {
		charSet[string(c)]++
	}
	if len(charSet) < 2 {
		return 0
	}

	for c, _ := range charSet {
		uniqueString = uniqueString + c
	}
	combinations := combinationsWithoutRepetitionsRecursive(uniqueString, len(charSet)-2, temp)
	for _, combination := range combinations {
		alternateString := ""
		for _, c := range s {
			if _, exists := combination[string(c)]; !exists {
				alternateString = alternateString + string(c)
			}
		}

		alternateStringLen := alternateLength(alternateString)
		if alternateStringLen > maxLen {
			maxLen = alternateStringLen
		}
	}

	return maxLen
}

func combinationsWithoutRepetitionsRecursive(s string, r int, temp map[string]int) []map[string]int {
	combinations := []map[string]int{}

	if len(temp) == r {
		newMap := map[string]int{}
		for key, v := range temp {
			newMap[key] = v
		}
		combinations = append(combinations, newMap)
	} else {
		for i, c := range s {
			temp[string(c)]++
			comb := combinationsWithoutRepetitionsRecursive(s[i+1:], r, temp)
			combinations = append(combinations, comb...)

			//pop last element

			delete(temp, string(c))
		}

	}

	return combinations
}

func combinatiosWithoutRepetitionIterative(s string, r int) []map[string]int {
	combinations := []map[string]int{}

	//develop
	for i := 0; i < len(s)-r; i++ {
		for k := i + 1; k < len(s)-r; k++ {

		}
	}

	return combinations
}

func alternateLength(s string) int32 {
	if len(s) <= 1 {
		return 0
	}

	var (
		last = ""
	)

	for _, c := range s {
		if string(c) == last {
			return 0
		}
		last = string(c)
	}

	return int32(len(s))
}
