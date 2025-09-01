package main

import "fmt"

func main() {

	result := beautifulBinaryString("0100101010")

	// 0100101010
	// 0110101010
	// 0110111010
	// 0110111011

	fmt.Println("result: ", result)

}

func beautifulBinaryString(b string) int32 {
	// Write your code here

	count := int32(0)
	binary := make([]string, len(b))
	for i, c := range b {
		binary[i] = string(c)
	}
	for i := 0; i < len(binary)-2; i++ {
		sub := binary[i] + binary[i+1] + binary[i+2]

		fmt.Println("sub: ", sub)

		if sub == "010" {
			binary[i+2] = "1"
			count++
		}
	}

	fmt.Println(binary)

	return count

}
