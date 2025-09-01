package main

import (
	"fmt"
	"strconv"
)

var (
	tests = []string{
		"90071992547409929007199254740993",
		"45035996273704964503599627370497",
		"22517998136852482251799813685249",
		"11258999068426241125899906842625",
		"562949953421312562949953421313",
		"90071992547409928007199254740993",
		"45035996273704963503599627370497",
		"22517998136852481251799813685249",
		"11258999068426240125899906842625",
		"562949953421312462949953421313",
	}
)

func main() {

	for _, v := range tests {
		separateNumbers(v)

	}

	// separateNumbers("45035996273704964503599627370497")

	// fmt.Println("Result: ", "")
}

func separateNumbers(s string) {
	// Write your code here
	if len(s) == 1 {
		fmt.Println("NO")
		return
	}

	isB, first := isBeautiful(s)
	if isB {
		fmt.Printf("YES %d\n", first)
	} else {
		fmt.Println("NO")
	}
}

func isBeautiful(s string) (bool, int) {
	first := 0
	half := s[:len(s)/2]
	for i, _ := range half {
		first, _ = strconv.Atoi(string(s[:i+1]))

		next := strconv.Itoa(first + 1)
		if string(s[i+1:i+1+len(next)]) == next {
			expected := buildExpectedBeautiful(first, len(s))
			if expected == s {
				return true, first
			}
		}
	}

	return false, 0
}

func buildExpectedBeautiful(first int, maxLen int) string {
	beautiful := strconv.Itoa(first)

	next := first + 1
	for len(beautiful) < maxLen {
		beautiful = beautiful + strconv.Itoa(next)
		next++
	}

	return beautiful
}
