package main

import "fmt"

func main() {
	result := stringConstruction("bccb")

	fmt.Println("Result: ", result)
}

func stringConstruction(s string) int32 {
	// Write your code here

	copyString := ""
	cost := int32(0)

	i := 0
	for i < len(s) {
		if copyString == "" {
			copyString = copyString + string(s[i])
			cost++
			i++

			continue
		}

		subs := ""
		k := i
		//find substring
		for _, c := range copyString {
			if k > len(s)-1 {
				break
			}
			if string(c) == string(s[k]) {
				subs = subs + string(c)
				k++
				continue
			}

			if subs == "" {
				continue
			}
			if subs != "" {
				break
			}

		}

		if subs != "" {
			copyString = copyString + subs
			i = i + len(subs)

			continue
		}

		copyString = copyString + string(s[i])
		cost++
		i++

	}

	if len(copyString) == len(s) {
		return cost
	}
	return -1
}
