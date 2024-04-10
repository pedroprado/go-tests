package main

import (
	"fmt"
	"time"
)

//longest common subsequence problem

func main() {
	// file, err := os.Open("input")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// reader := bufio.NewReader(file)
	// buffer := make([]byte, 1000)

	// var s string

	// for {
	// 	n, err := reader.Read(buffer)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatal(err)
	// 	}

	// 	readString := buffer[0:n]
	// 	readString = []byte(strings.TrimSpace(string(readString)))
	// 	s = s + string(readString)
	// }

	// n := int32(len(s))
	// n := int32(10)

	s1 := "WEWOUCUIDGCGTRMEZEPXZFEJWISRSBBSYXAYDFEJJDLEBVHHKS"
	s2 := "FDAGCXGKCTKWNECHMRXZWMLRYUCOCZHJRRJBOAJOQJZZVUYXIC"
	start := time.Now()
	fmt.Println("Starting process")
	result := commonChild3(s1, s2)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)
}

func commonChild(s1 string, s2 string) int32 {
	// Write your code here

	subs1 := ""
	index2 := 0

	for _, l := range s1 {
		letter1 := string(l)
		letter2 := string(s2[index2])

		if letter1 != letter2 {
			continue
		}
		if letter1 == letter2 {

		}
	}

	subs2 := ""
	index1 := 0
	for _, l := range s2 {
		letter := string(l)
		if letter == string(s1[index1]) {
			subs2 += letter
		}
	}

	if len(subs1) > len(subs2) {
		return int32(len(subs1))
	}

	return int32(len(subs2))
}

func commonChild2(s1 string, s2 string) int32 {
	// Write your code here
	indexMap1 := getIndexMap(s1)
	indexMap2 := getIndexMap(s2)

	subs := []string{}
	for i := 0; i < len(s1); i++ {
		sub := s1[i:]

		subsFinal1 := ""
		lastIndex2 := 0
		for _, l1 := range sub {
			letter1 := string(l1)
			indexes := indexMap2[letter1]
			if len(indexes) == 0 {
				continue
			}

			foundIndex := false
			for _, index2 := range indexes {
				if index2 >= lastIndex2 {
					//found the index where letter2 = letter1
					//add the letter to child string
					subsFinal1 += letter1
					foundIndex = true
					lastIndex2 = index2
					break
				}
			}

			if lastIndex2 == len(s1)-1 {
				break
			}
			//remove the index of that letter2 from the indexes array
			if foundIndex {
				newIndexes := []int{}
				for _, v := range indexes {
					if v != lastIndex2 {
						newIndexes = append(newIndexes, v)
					}
				}

				indexMap2[letter1] = newIndexes
			}
		}

		if subsFinal1 != "" {
			subs = append(subs, subsFinal1)
		}
	}

	for j := 0; j < len(s2); j++ {
		sub := s2[j:]

		subsFinal2 := ""
		lastIndex1 := 0
		for _, l2 := range sub {
			letter2 := string(l2)
			indexes := indexMap1[letter2]
			if len(indexes) == 0 {
				continue
			}

			foundIndex := false
			for _, index1 := range indexes {
				if index1 >= lastIndex1 {
					subsFinal2 += letter2
					foundIndex = true
					lastIndex1 = index1
					break
				}
			}
			if lastIndex1 == len(s2)-1 {
				break
			}

			//remove the index of that letter2 from the indexes array
			if foundIndex {
				newIndexes := []int{}
				for _, v := range indexes {
					if v != lastIndex1 {
						newIndexes = append(newIndexes, v)
					}
				}

				indexMap1[letter2] = newIndexes
			}
		}

		if subsFinal2 != "" {
			subs = append(subs, subsFinal2)
		}
	}

	longest := getLongest(subs)

	return int32(len(longest))
}

func getIndexMap(s string) map[string][]int {
	indexMap := map[string][]int{}
	for i, l := range s {
		letter := string(l)
		if len(indexMap[letter]) == 0 {
			indexMap[letter] = []int{i}
		} else {
			indexMap[letter] = append(indexMap[letter], i)
		}
	}

	return indexMap
}

func getLongest(substrings []string) string {
	fmt.Println(substrings)
	longest := 0
	sub := ""

	for _, s := range substrings {
		if len([]rune(s)) >= longest {
			longest = len(s)
			sub = s
		}
	}

	return sub
}

// compare
func commonChild3(s1 string, s2 string) int32 {
	// Longest common subsequence algorithm

	table := make([][]int, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		table[i] = make([]int, len(s2)+1)
	}

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				table[i][j] = 0
				continue
			}
			if string(s1[i-1]) == string(s2[j-1]) {
				table[i][j] = table[i-1][j-1] + 1
				continue

			}

			table[i][j] = max(table[i-1][j], table[i][j-1])
		}
	}

	return int32(table[len(s1)][len(s2)])
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}

	return v2
}

// subsequences = characters DO NOT need to occupy CONSECUTIVE positions
func subsequences(s string) []string {
	subs := []string{}

	return subs
}

// substrings = characters need to occupy CONSECUTIVE positions
func substrings(s string) []string {
	subs := []string{}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			sub := s[i:j]

			subs = append(subs, sub)
		}
	}
	return subs
}
