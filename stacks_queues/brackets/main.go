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
	brackets := "(){}[()[][]]{}(())()[[([])][()]{}{}(({}[]()))()[()[{()}]][]]"

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

	openingBrackets := []string{firstBracket}

	for i := 1; i < len(s); i++ {
		nextBracket := string(s[i])
		if isOpeningBracket(nextBracket) {
			openingBrackets = append(openingBrackets, nextBracket)
			continue
		}

		//is a closing bracket
		if len(openingBrackets) == 0 {
			return "NO"
		}
		openingBracket, exists := closingBracketsMap[nextBracket]
		if !exists {
			return "NO"
		}
		if openingBracket != openingBrackets[len(openingBrackets)-1] {
			return "NO"
		}

		if len(openingBrackets) == 1 {
			openingBrackets = []string{}
			continue
		}
		// pop
		newOpeningBrackets := []string{}
		newOpeningBrackets = append(newOpeningBrackets, openingBrackets[:len(openingBrackets)-1]...)
		openingBrackets = newOpeningBrackets

	}

	return "YES"
}

func isOpeningBracket(c string) bool {
	_, exists := openingBracketsMap[c]
	return exists
}

func isClosingBracket(c string) bool {
	_, exists := closingBracketsMap[c]
	return exists
}
