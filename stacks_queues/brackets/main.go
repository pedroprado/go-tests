package main

import (
	"fmt"
	"time"
)

var (
	openingBracketsMap = map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	closingBracketsMap = map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
)

func main() {
	brackets := "{[()]}"

	start := time.Now()
	fmt.Println("Starting process")
	result := isBalanced(brackets)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println("Result: ", result)
}

func isBalanced(s string) string {
	if len(s) == 0 {
		return "NO"
	}
	if len(s)%2 != 0 {
		return "NO"
	}
	firstBracket := string(s[0])
	if !isOpeningBracket(firstBracket) {
		return "NO"
	}

	brackets := []string{}

	for i := 0; i < len(s); i++ {
		bracket := string(s[i])

		if isOpeningBracket(bracket) {
			brackets = append(brackets, bracket)

			continue
		}

		if isClosingBracket(bracket) && lastMatchesClosingBacket(bracket, brackets) {
			brackets, _ = popBracket(brackets)
			newBrackets := []string{}
			newBrackets = append(newBrackets, brackets[0:len(brackets)]...)
			brackets = newBrackets

			continue
		}

		return "NO"
	}

	if len(brackets) == 0 {
		return "YES"
	}

	return "NO"
}

func isOpeningBracket(bracket string) bool {
	_, exists := openingBracketsMap[bracket]
	return exists
}

func isClosingBracket(bracket string) bool {
	_, exists := closingBracketsMap[bracket]
	return exists
}

func pushBracket(bracket string, brackets []string) []string {
	brackets = append(brackets, bracket)
	return brackets
}

func popBracket(brackets []string) ([]string, string) {
	newBrackets := []string{}
	newBrackets = append(newBrackets, brackets[0:len(brackets)-1]...)

	return newBrackets, brackets[len(brackets)-1]
}

func lastMatchesClosingBacket(closingBracket string, brackets []string) bool {
	if len(brackets) == 0 {
		return false
	}

	openingBracket, _ := closingBracketsMap[closingBracket]

	return openingBracket == brackets[len(brackets)-1]
}
