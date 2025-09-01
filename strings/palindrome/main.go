package main

import "fmt"

func main() {
	result := theLoveLetterMystery("abcba")

	fmt.Println("result: ", result)
}

func palindrome(s string) bool {

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

func theLoveLetterMystery(s string) int32 {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetArr := make([]string, len(alphabet))
	alphabetIndex := map[string]int{}
	changes := int32(0)

	for i, c := range alphabet {
		alphabetArr[i] = string(c)
		alphabetIndex[string(c)] = i
	}

	if len(s) == 0 {
		return 0
	}

	i := 0
	for i <= len(s)-1-i {
		left := string(s[i])
		right := string(s[len(s)-1-i])
		fmt.Println(i)

		if left > right {
			for left > right {
				leftIndex, _ := alphabetIndex[left]
				leftIndex--
				left = string(alphabet[leftIndex])
				changes++

			}
			i++
			continue
		}
		if right > left {
			for right > left {
				rightIndex, _ := alphabetIndex[right]
				rightIndex--
				right = string(alphabet[rightIndex])
				changes++
			}
			i++
			continue
		}
		i++

	}

	return changes
}
