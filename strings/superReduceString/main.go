package main

import "fmt"

func main() {

	result := superReducedString("baab")

	fmt.Println("result: ", result)
}

func superReducedString(s string) string {
	// Write your code here

	i := 0
	for i < len(s) {
		fmt.Println(i)
		if len(s) < 2 {
			fmt.Println("stoped. final string: ", s)
			break
		}

		if i < len(s)-1 {
			if s[i] == s[i+1] {
				fmt.Println("deleting from string: ", s)
				left := ""
				right := ""

				if i > 0 {
					for j := 0; j < i; j++ {
						left = left + string(s[j])
					}
				}

				if len(s[i+1:]) > 1 {
					for k := i + 2; k < len(s); k++ {
						right = right + string(s[k])
					}
				}

				s = left + right
				i = 0
				continue
			}
		}

		i++

	}

	if s == "" {
		return "Empty String"
	}
	return s
}
