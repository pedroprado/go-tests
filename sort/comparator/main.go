package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("out")
	if err != nil {
		panic(err)
	}

	defer stdout.Close()

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	if err != nil {
		panic(err)
	}
	q := int(qTemp)

	players := make(Players, q)

	for i := 0; i < q; i++ {
		s := readLine(reader)
		line := strings.Split(s, " ")

		name := line[0]
		score := line[1]
		scoreNum, err := strconv.Atoi(score)
		if err != nil {
			panic(err)
		}

		players[i] = Player{Name: name, Score: scoreNum}

	}

	sort.Sort(players)

	for _, player := range players {
		fmt.Printf("%s %d\n", player.Name, player.Score)

	}

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

type Player struct {
	Name  string
	Score int
}

type Players []Player

func (players Players) Len() int {
	return len(players)
}

func (players Players) Less(i, j int) bool {
	v := comparePlayers(players[i], players[j])
	return v == -1
}

func (players Players) Swap(i, j int) {
	players[i], players[j] = players[j], players[i]
}

func comparePlayers(player1, player2 Player) int {
	if player1.Score > player2.Score {
		return -1
	}
	if player1.Score == player2.Score {
		if player1.Name > player2.Name {
			return 1
		}
		if player1.Name == player2.Name {
			return 0
		}

		return -1
	}

	return 1
}

type PlayerChecker struct{}

func (ref PlayerChecker) Compare(player1, player2 Player) int {
	if player1.Score > player2.Score {
		return 1
	}
	if player1.Score == player2.Score {
		if player1.Name > player2.Name {
			return 1
		}
		if player1.Name == player2.Name {
			return 0
		}

		return -1
	}

	return -1
}
