package main

func main() {
	// prices := []int32{1, 10, 29, 4, 9, 3, 11}

}

func maximumToys(prices []int32, k int32) int32 {

	bubbleSort(prices)

	priceSum := int32(0)
	numToys := int32(0)

	for _, price := range prices {
		if priceSum+price > k {
			break
		}

		priceSum = priceSum + price
		numToys++

	}

	return numToys
}

func bubbleSort(arr []int32) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
