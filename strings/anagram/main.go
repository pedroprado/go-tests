package main

import "fmt"

func main() {

	result := makingAnagrams("cde", "abc")

	fmt.Println("result: ", result)
}

func anagram(s string) int32 {
	// Write your code here
	if len(s)%2 != 0 {
		return -1
	}
	middle := len(s) / 2
	left := s[0:middle]
	right := s[middle:]
	fmt.Printf("%+v\n", left)
	fmt.Printf("%+v\n", right)

	freqLeft := map[string]int{}
	freqRight := map[string]int{}
	for _, c := range left {
		freqLeft[string(c)]++
	}
	for _, c := range right {
		freqRight[string(c)]++
	}
	fmt.Printf("%+v\n", freqLeft)
	fmt.Printf("%+v\n", freqRight)

	numberChanges := int32(0)
	for letter, freqR := range freqRight {
		freqL, exists := freqLeft[letter]
		if exists && (freqR == freqL || freqR < freqL) {
			continue
		}
		numberChanges = numberChanges + int32(freqR-freqL)
	}

	return numberChanges
}

func makingAnagrams(s1 string, s2 string) int32 {
	freqS1 := map[string]int{}
	freqS2 := map[string]int{}
	for _, c := range s1 {
		freqS1[string(c)]++
	}
	for _, c := range s2 {
		freqS2[string(c)]++
	}
	fmt.Printf("%+v\n", freqS1)
	fmt.Printf("%+v\n", freqS2)

	numberDeletions := int32(0)
	for letter1, freq1 := range freqS1 {
		freq2, exists := freqS2[letter1]
		if !exists {
			numberDeletions = numberDeletions + int32(freq1)
			delete(freqS1, letter1)
			continue
		}
		if freq1 > freq2 {
			numberDeletions = numberDeletions + int32(freq1-freq2)
			freqS1[letter1] = freqS1[letter1] - (freq1 - freq2)
			continue
		}
		if freq2 > freq1 {
			numberDeletions = numberDeletions + int32(freq2-freq1)
			freqS2[letter1] = freqS2[letter1] - (freq2 - freq1)
			continue
		}
	}

	for letter2, freq2 := range freqS2 {
		_, exists := freqS1[letter2]
		if !exists {
			numberDeletions = numberDeletions + int32(freq2)
			delete(freqS2, letter2)
		}
	}

	fmt.Println("final")
	fmt.Printf("%+v\n", freqS1)
	fmt.Printf("%+v\n", freqS2)

	return numberDeletions
}
