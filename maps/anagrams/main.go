package main

import (
	"fmt"
	"sort"
	"time"
)

//number of substrings = n*(n+1)/2

func main() {

	s1 := "ofeqjnqnxwidhbuxxhfwargwkikjqwyghpsygjxyrarcoacwnhxyqlrviikfuiuotifznqmzpjrxycnqktkryutpqvbgbgthfges"
	// s2 := "listen"

	start := time.Now()
	result := sherlockAndAnagrams2(s1)
	// result := sherlockAndAnagrams(s1)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)

}

func sherlockAndAnagrams2(s string) int32 {
	substringsMapFrequency := map[string]int{}

	var count int32

	for size := 1; size <= len(s)-1; size++ {
		for j := 0; j <= len(s)-size; j++ {
			word := s[j : j+size]
			r := []rune(word)
			sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })

			sorted := string(r)

			substringsMapFrequency[sorted] = substringsMapFrequency[sorted] + 1

		}
	}

	for _, frequency := range substringsMapFrequency {

		//number of anagrams = number os combinations without repetition of the N times "XX"
		combinations := combinationWithoutRepetition(frequency)

		count += int32(combinations)
	}

	return count
}

// combination without repetition for (A,B,C,B,E) (size 5) = 4+3+2+1 = (size-1)+(size-2)+...+(1)
func combinationWithoutRepetition(m int) int {
	total := 0

	for i := m - 1; i > 0; i-- {
		total += i
	}
	return total
}

// Count anagrams for all combinations of strings by size
// O(n2) time complexity
func sherlockAndAnagrams(s string) int32 {
	substrings := getSubstrings(s)

	var count int32
	for i := 1; i <= len(s)-1; i++ {
		numberAnagrams := countAnagrams(substrings[i])
		count = count + numberAnagrams
	}

	return count
}

func getSubstrings(s string) map[int][]string {
	substringsMap := map[int][]string{}

	for size := 1; size <= 1; size++ {

		substrings := []string{}
		for j := 0; j <= len(s)-size; j++ {
			word := s[j : j+size]
			substrings = append(substrings, word)
		}

		substringsMap[size] = substrings
	}

	return substringsMap
}

// Count anagrams for all combinations of strings in arr
// O(n2) time complexity
func countAnagrams(arr []string) int32 {
	var count int32

	for i := 0; i < len(arr)-1; i++ {
		first := arr[i]

		for j := i + 1; j < len(arr); j++ {
			second := arr[j]
			isAnagram := anagram(first, second)
			if isAnagram {
				count++
			}
		}
	}

	return count
}

// this compares if two strings are anagrams
// O(n) time complexity (3n)
func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Map := map[string]int{}
	s2Map := map[string]int{}

	for i := range s1 {
		letter1 := string(s1[i])
		s1Map[string(letter1)] = s1Map[string(letter1)] + 1

		letter2 := string(s2[i])
		s2Map[string(letter2)] = s2Map[string(letter2)] + 1
	}

	for letter1, count1 := range s1Map {
		count2, exists := s2Map[letter1]
		if !exists || count1 != count2 {

			return false
		}
	}

	return true
}
