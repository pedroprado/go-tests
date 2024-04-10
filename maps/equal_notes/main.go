package main

import (
	"fmt"
)

func main() {
	magazine := []string{"Atack", "on", "titan"}
	note := []string{"Atack", "on", "titan"}

	checkMagazine(magazine, note)

}

func checkMagazine(magazine []string, note []string) {
	mapMagazine := map[string]int{}

	if len(magazine) == 0 && len(note) != 0 {
		fmt.Println("No")
		return
	}

	for _, word := range magazine {
		_, exists := mapMagazine[word]
		if !exists {
			mapMagazine[word] = 1

		} else {
			mapMagazine[word] = mapMagazine[word] + 1
		}
	}

	for _, word := range note {
		count, exists := mapMagazine[word]
		if !exists || count == 0 {
			fmt.Println("No")
			return
		}
		mapMagazine[word] = mapMagazine[word] - 1
	}
	fmt.Println("Yes")
	return
}
