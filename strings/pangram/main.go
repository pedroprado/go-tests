package main

import (
	"fmt"
	"strings"
)

func main() {
	result := pangrams("We promptly judged antique ivory buckles for the next prize")

	fmt.Println("Result: ", result)
}

func pangrams(s string) string {
	// Write your code here
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	lowerMap := map[string]int{}

	count := 0
	for _, c := range alphabet {
		lowerMap[string(c)]++
	}

	for _, c := range s {
		letter := strings.ToLower(string(c))
		if _, exists := lowerMap[letter]; exists {
			count++
			delete(lowerMap, letter)
			continue
		}

	}

	fmt.Println(count)
	if count == 26 {
		return "pangram"
	}

	return "not pangram"
}
