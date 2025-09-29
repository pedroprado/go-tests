package main

import (
	"fmt"
	"sort"
)

func main() {
	sticks := []int32{1, 1, 1, 3, 3}

	result := maximumPerimeterTriangle(sticks)

	fmt.Println("Result: ", result)
}

type ArrInt []int32

func (arr ArrInt) Len() int {
	return len(arr)
}

func (arr ArrInt) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr ArrInt) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func maximumPerimeterTriangle(sticks []int32) []int32 {
	// Write your code here
	invalidResponse := []int32{-1}
	if len(sticks) < 3 {
		return invalidResponse
	}

	sort.Sort(ArrInt(sticks))

	validTriangles := [][]int32{}
	for i := 0; i < len(sticks)-2; i++ {
		sides := []int32{sticks[i], sticks[i+1], sticks[i+2]}
		if isValidTriangle(sides[0], sides[1], sides[2]) {
			validTriangles = append(validTriangles, sides)
		}
	}

	fmt.Println(validTriangles)

	if len(validTriangles) == 0 {
		return invalidResponse
	}

	return greeedyChoice(validTriangles)
}

func isValidTriangle(a, b, c int32) bool {
	return isInequalTriangle(a, b, c) || isEquilateralTriange(a, b, c)
}

func isInequalTriangle(a, b, c int32) bool {
	return a+b > c && a+c > b && b+c > a
}

func isEquilateralTriange(a, b, c int32) bool {
	return a == b && b == c
}

// greedy choice : maximum perimeter = maximun sides
func greeedyChoice(triangles [][]int32) []int32 {
	longestMaximumSide := int32(-1)
	longestMaximumSideCandidates := []int{}
	for i, t := range triangles {
		if t[2] > longestMaximumSide {
			longestMaximumSide = t[2]
			if len(longestMaximumSideCandidates) == 0 {
				longestMaximumSideCandidates = append(longestMaximumSideCandidates, i)
				continue
			}
			longestMaximumSideCandidates[len(longestMaximumSideCandidates)-1] = i
			continue
		}
		if t[2] == longestMaximumSide {
			longestMaximumSideCandidates = append(longestMaximumSideCandidates, i)
		}
	}

	fmt.Println(longestMaximumSideCandidates)

	if len(longestMaximumSideCandidates) == 1 {
		return triangles[longestMaximumSideCandidates[0]]
	}

	longestMinimumSide := int32(-1)
	longestMinimumSideCandidates := []int{}
	for _, indexTriangle := range longestMaximumSideCandidates {
		t := triangles[indexTriangle]
		if t[0] > longestMinimumSide {
			longestMinimumSide = t[0]
			if len(longestMinimumSideCandidates) == 0 {
				longestMinimumSideCandidates = append(longestMinimumSideCandidates, indexTriangle)
				continue
			}
			longestMaximumSideCandidates[len(longestMinimumSideCandidates)-1] = indexTriangle
			continue
		}
		if t[0] == longestMinimumSide {
			longestMinimumSideCandidates = append(longestMinimumSideCandidates, indexTriangle)
		}
	}

	fmt.Println(longestMinimumSideCandidates)

	if len(longestMinimumSideCandidates) == 1 {
		return triangles[longestMinimumSideCandidates[0]]
	}

	return triangles[longestMinimumSideCandidates[0]]
}
