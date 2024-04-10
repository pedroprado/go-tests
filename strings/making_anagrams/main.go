package main

import "fmt"

func main() {

	var str1 = "fcrxzwscanmligyxyvym"
	var str2 = "jxwtrhvujlmrpdoqbisbwhmgpmeoke"

	numberOfDeletions := makeAnagram(str1, str2)

	fmt.Println("number of deletions: ", numberOfDeletions)

}

func makeAnagram(a string, b string) int32 {
	// Write your code here

	var count int32
	freqMapA := map[string]int{}
	freqMapB := map[string]int{}

	for i := range a {
		letterA := string(a[i])
		freqMapA[letterA]++
	}

	for i := range b {
		letterB := string(b[i])
		freqMapB[letterB]++
	}

	for letterA, freqA := range freqMapA {
		freqB, _ := freqMapB[letterA]

		if freqA > freqB {
			deletions := freqA - freqB
			freqMapA[letterA] = freqMapA[letterA] - deletions
			count = count + int32(deletions)
		}
		if freqB > freqA {
			deletions := freqB - freqA
			freqMapB[letterA] = freqMapB[letterA] - deletions
			count = count + int32(deletions)
		}
	}

	for letterB, freqB := range freqMapB {
		freqA, _ := freqMapA[letterB]

		if freqA > freqB {
			deletions := freqA - freqB
			freqMapA[letterB] = freqMapA[letterB] - deletions
			count = count + int32(deletions)
		}
		if freqB > freqA {
			deletions := freqB - freqA
			freqMapB[letterB] = freqMapB[letterB] - deletions
			count = count + int32(deletions)
		}
	}

	return count
}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	freqMapA := map[string]int{}
	freqMapB := map[string]int{}
	for i := range a {
		letterA := string(a[i])
		freqMapA[letterA]++

		letterB := string(b[i])
		freqMapB[letterB]++
	}

	for letterA, freqA := range freqMapA {
		freqB, exists := freqMapB[letterA]
		if !exists || freqB != freqA {
			return false
		}
	}

	return true
}
