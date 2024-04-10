package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("teste1")
	checkError(err)

	reader := bufio.NewReaderSize(file, 16*1024*1024)

	stdout, err := os.Create("out")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	start := time.Now()
	ans := freqQuery2(queries)
	fmt.Println("Time to process(ms): ", time.Since(start).Milliseconds())
	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func freqQuery(queries [][]int32) []int32 {
	result := []int32{}
	freqMap := map[int32]int{}

	for _, v := range queries {
		if v[0] == 1 {
			//add

			freqMap[v[1]] = freqMap[v[1]] + 1

		}
		if v[0] == 2 {
			//remove

			if freqMap[v[1]] > 0 {
				freqMap[v[1]] -= 1
			}

		}
		if v[0] == 3 {
			//print
			expectedFreq := v[1]

			found := false

			for _, frequency := range freqMap {
				if frequency == int(expectedFreq) {
					found = true
					break
				}
			}

			if found {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
	}

	return result
}

func freqQuery2(queries [][]int32) []int32 {
	result := []int32{}
	freqMap := map[int32]int{}

	freqFreqMap := map[int]int{}

	for _, v := range queries {
		if v[0] == 1 {
			//add
			if freqFreqMap[freqMap[v[1]]] > 0 {
				freqFreqMap[freqMap[v[1]]] -= 1
			}

			freqMap[v[1]] = freqMap[v[1]] + 1

			freqFreqMap[freqMap[v[1]]] += 1
		}
		if v[0] == 2 {
			//remove
			if freqFreqMap[freqMap[v[1]]] > 0 {
				freqFreqMap[freqMap[v[1]]] -= 1
			}

			if freqMap[v[1]] > 0 {
				freqMap[v[1]] -= 1
			}

			freqFreqMap[freqMap[v[1]]] += 1

		}
		if v[0] == 3 {
			//print
			expectedFreq := v[1]

			num := freqFreqMap[int(expectedFreq)]

			if num > 0 {
				result = append(result, 1)
			} else {
				result = append(result, 0)
			}
		}
	}

	return result
}
