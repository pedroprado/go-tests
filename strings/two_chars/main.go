package main

import "fmt"

func main() {

	result := alternate("beabeefeab")

	fmt.Println("result: ", result)
}

func alternate(s string) int32 {
	// Write your code here
	mapIndexes := map[string][]int{}

	for i, c := range s {
		if arr := mapIndexes[string(c)]; len(arr) == 0 {
			mapIndexes[string(c)] = []int{i}
		} else {
			mapIndexes[string(c)] = append(mapIndexes[string(c)], i)
		}
	}

	longest := int32(0)
	for c, _ := range mapIndexes {
		newString := ""
		for _, letter := range s {
			if string(letter) != c {
				newString = newString + string(letter)
			}
		}

	}

	return longest

}

func calculateLentgh(s string) int32 {
	last := string(s[0])
	alternated := last

	for i, c := range s {
		nextChar := string(c)
		if i == 0 {
			continue
		}

		if nextChar != last {
			alternated = alternated + nextChar
			last = nextChar
		} else {
			sLen := len(s)
			alternateLen := len(alternated)
			if sLen%alternateLen != 0 {
				return 0
			}
			numbSubstrings := sLen / alternateLen
			j := 0
			count := 0
			for count < numbSubstrings {
				sub := s[j:alternateLen]
				if sub != alternated {
					return 0
				}
				j = j + alternateLen
			}

		}
	}

	return int32(len(alternated))
}
