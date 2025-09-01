package main

import "fmt"

var (
	tests = []string{
		"acxz",
	}
)

func main() {

	result := funnyString("acxz")

	fmt.Println("Result: ", result)
}

func funnyString(s string) string {
	// Write your code here
	reversed := ""
	for _, c := range s {
		reversed = string(c) + reversed
	}

	for i := 0; i < len(s)-1; i++ {
		absNormal := absoluteDifference(int32(s[i]), int32(s[i+1]))
		absReversed := absoluteDifference(int32(reversed[i]), int32(reversed[i+1]))

		if absNormal != absReversed {
			return "Not Funny"
		}
	}

	return "Funny"

}

func absoluteDifference(x, y int32) int32 {
	if x > y {
		return x - y
	}
	return y - x
}
