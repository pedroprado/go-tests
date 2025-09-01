package main

import (
	"fmt"
	"strings"
)

func main() {

	result := caesarCipher("abcdefghijklmnopqrstuvwxyz", 2)

	fmt.Println("Result: ", result)
}

func caesarCipher(s string, k int32) string {
	// Write your code here
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetIndexLower := map[string]*int{}
	alphabetIndexUpper := map[string]*int{}
	originalAlphabetLower := []string{}

	for i, c := range alphabet {
		originalAlphabetLower = append(originalAlphabetLower, string(c))
		index := i
		alphabetIndexLower[string(c)] = &index
		alphabetIndexUpper[strings.ToUpper(string(c))] = &index
	}

	rotatedLower := rotateLeft(originalAlphabetLower, k)

	rotatedString := ""
	for _, c := range s {
		charIndex := alphabetIndexLower[string(c)]
		if charIndex != nil {
			rotatedCharLower := rotatedLower[*charIndex]
			rotatedString = rotatedString + rotatedCharLower
			continue
		}

		charIndex = alphabetIndexUpper[string(c)]
		if charIndex != nil {
			rotatedCharLower := rotatedLower[*charIndex]
			rotatedString = rotatedString + strings.ToUpper(rotatedCharLower)
			continue
		}

		rotatedString = rotatedString + string(c)
	}

	return rotatedString
}

func rotateLeft(original []string, k int32) []string {
	rotated := []string{}

	k = k % int32(len(original))

	rotated = append(rotated, original[k:]...)
	rotated = append(rotated, original[:k]...)

	return rotated
}
