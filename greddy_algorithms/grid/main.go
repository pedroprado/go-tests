package main

import "fmt"

func main() {
	arr := []string{"ppp", "ypp", "wyw"}
	result := gridChallenge(arr)

	fmt.Println("Result : ", result)
}

func gridChallenge(grid []string) string {
	// Write your code here

	n := len(grid)
	if n < 2 {
		return "YES"
	}
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	alphabetIndex := map[string]int{}
	for i, c := range alphabet {
		alphabetIndex[string(c)] = i
	}

	// sorting rows ascending
	for i, row := range grid {
		freqArr := make([]int, 26)
		for _, letter := range row {
			//convert letter to index
			i, _ := alphabetIndex[string(letter)]
			freqArr[i]++
		}

		sortedRow := ""
		for j, freq := range freqArr {
			if freq > 0 {
				// convert index to letter
				letter := string(alphabet[j])
				for range freq {
					sortedRow = sortedRow + letter
				}
			}
		}

		grid[i] = sortedRow

	}

	for i := range n {
		previous := string(grid[0][i])
		iPrevious, _ := alphabetIndex[previous]
		for j := 1; j < n; j++ {
			next := string(grid[j][i])
			iNext, _ := alphabetIndex[next]
			if iPrevious <= iNext {
				previous = next
				continue
			}
			return "NO"

		}
	}

	return "YES"

}
