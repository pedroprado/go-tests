package main

import "fmt"

func main() {
	result := palindromeIndex("hgygsvlfcwnswtuhmyaljkqlqjjqlqkjlaymhutwsnwcwflvsgygh")

	fmt.Println("result: ", result)
}

func palindromeIndex(s string) int32 {
	if len(s) == 0 {
		return -1
	}

	i := 0
	for i <= len(s)-1-i {
		leftIndex := i
		rightIndex := len(s) - 1 - i
		left := string(s[leftIndex])
		right := string(s[rightIndex])
		if left != right {
			if palindrome(s[leftIndex:rightIndex]) {
				return int32(rightIndex)
			}
			if palindrome(s[leftIndex+1 : rightIndex+1]) {
				return int32(leftIndex)
			}

			return -1
		}
		i++
	}

	return -1
}

func palindrome(s string) bool {
	fmt.Println("testing: ", s)

	if len(s) == 0 {
		return false
	}

	i := 0
	for i <= len(s)-1-i {
		left := string(s[i])
		right := string(s[len(s)-1-i])
		if left != right {
			return false
		}
		i++
	}

	return true
}
