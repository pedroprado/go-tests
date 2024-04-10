package main

import (
	"fmt"
	"time"
)

func main() {
	s1 := "ABABABAB"

	start := time.Now()
	result := alternatingCharacters(s1)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)
}

func alternatingCharacters(s string) int32 {
	// Write your code here

	letters := make([]string, len(s))

	for i, v := range s {
		letter := string(v)
		letters[i] = letter
	}

	var deletions int32

	lastLetter := string(s[0])
	for i := 1; i < len(letters); i++ {
		currentLetter := string(s[i])

		if lastLetter == currentLetter {
			lastLetter = currentLetter
			deletions++
			continue
		}

		if lastLetter != currentLetter {
			lastLetter = currentLetter
			continue

		}

	}

	return deletions
}
