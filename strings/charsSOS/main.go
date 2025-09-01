package main

import "fmt"

func main() {

	result := marsExploration("SOSSPSSQSSOR")

	fmt.Println("Result: ", result)
}

func marsExploration(s string) int32 {
	// Write your code here
	count := int32(0)
	for i, c := range s {
		remainder := (i + 1) % 3
		fmt.Println(remainder)
		if remainder == 1 && string(c) == "S" {
			continue
		}
		if remainder == 2 && string(c) == "O" {
			continue
		}
		if remainder == 0 && string(c) == "S" {
			continue
		}
		count++
		fmt.Println("count: ", count)

	}

	return count
}
