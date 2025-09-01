package main

import "fmt"

func main() {

	result := alternate("beabeefeab")

	fmt.Println("result: ", result)
}

func alternate(s string) int32 {
	// Write your code here
	var (
		max1    = 0
		keyMax1 = ""
		max2    = 0
		keyMax2 = ""

		freq = map[string]int{}
	)

	for _, c := range s {
		freq[string(c)]++
	}

	fmt.Println(freq)
	if len(freq) <= 2 {
		return int32(0)
	}

	for key, f := range freq {
		if f > max1 {
			max2 = max1
			keyMax2 = keyMax1

			max1 = f
			keyMax1 = key
		}
	}

	if len(freq) == 3 {
		return int32(max1)
	}

	delete(freq, keyMax1)
	delete(freq, keyMax2)

	fmt.Println(max1)
	fmt.Println(max2)

	return int32(max1 + max2)
}
