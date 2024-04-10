package main

import (
	"fmt"
	"time"
)

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
	n := int32(7)
	s := "abcbaba"
	fmt.Println("string size: ", n)
	start := time.Now()
	fmt.Println("Starting process")
	result := substrCount2(n, s)

	fmt.Println("Time to complete: ", time.Since(start))
	fmt.Println(result)
}

func substrCount(n int32, s string) int64 {
	var count int64
	for i := 0; i <= len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			subs := s[i:j]

			// evens cannot be special string
			if len(subs)%2 == 0 {
				continue
			}
			isSpecial := specialString(subs)
			if isSpecial {
				count++
			}
		}
	}

	return count
}

func specialString(s string) bool {
	freqMap := map[string]int{}
	for _, l := range s {
		letter := string(l)
		freqMap[letter]++
	}

	if len(freqMap) == 1 {
		return true
	}

	if len(s)%2 != 0 && len(freqMap) == 2 {
		freq1Letter := ""
		for letter, freq := range freqMap {
			if freq == 1 {
				freq1Letter = letter
			}
		}

		if freq1Letter == "" {
			return false
		}

		mid := len(s) / 2
		if string(s[mid]) == freq1Letter {
			return true
		}
	}

	return false
}

func substrCount2(n int32, s string) int64 {
	var count int64
	count += int64(len(s))
	for i := 0; i <= len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				count++
				continue
			}

			firstHalf := string(s[i:j])
			length := len(firstHalf)

			nextIndex := j + 1
			lastindex := nextIndex + length

			if lastindex <= len(s) {
				secondHalf := s[nextIndex:lastindex]
				if firstHalf == secondHalf {

					count++
				}
			}

			break
		}
	}

	return count
}
