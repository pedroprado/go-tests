package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := "abbac"

	start := time.Now()
	result := isValid(s1)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)
}

func isValid(s string) string {
	// Write your code here

	freqMap := map[string]int{}
	freqMap2 := map[int][]string{}
	for _, l := range s {
		letter := string(l)
		freqMap[letter]++
	}

	for letter, freq := range freqMap {
		if freqMap2[freq] == nil {
			freqMap2[freq] = []string{letter}
		} else {
			freqMap2[freq] = append(freqMap2[freq], letter)
		}
	}

	fmt.Println(freqMap2)

	if len(freqMap2) == 1 {
		return "YES"
	}

	if len(freqMap2) == 2 {
		freqs := []int{}
		for freq := range freqMap2 {
			freqs = append(freqs, freq)
		}

		freq1 := freqs[0]
		freq2 := freqs[1]

		letters1 := freqMap2[freq1]
		letters2 := freqMap2[freq2]

		if (freq1 == 1 && len(letters1) == 1) || (freq2 == 1 && len(letters2) == 1) {
			return "YES"
		}

		consecutive, higher := areConsecutive(freq1, freq2)
		if !consecutive {
			return "NO"
		}

		lettersHigher := freqMap2[higher]
		if len(lettersHigher) == 1 {
			return "YES"
		}
	}

	return "NO"
}

func areConsecutive(freq1, freq2 int) (bool, int) {
	if freq1 == freq2+1 {
		return true, freq1
	}
	if freq2 == freq1+1 {
		return true, freq2
	}

	return false, 0
}
