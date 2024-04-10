package main

import "fmt"

func main() {
	s1 := "asxz"
	s2 := "092oz"

	result := twoStrings(s1, s2)

	fmt.Println(result)
}

func twoStrings(s1, s2 string) string {
	s1Map := map[string]int{}
	s2Map := map[string]int{}

	for _, v1 := range s1 {
		_, exists := s1Map[string(v1)]
		if !exists {
			s1Map[string(v1)] = 1
		} else {
			s1Map[string(v1)] = s1Map[string(v1)] + 1
		}
	}

	for _, v2 := range s2 {
		_, exists := s2Map[string(v2)]
		if !exists {
			s2Map[string(v2)] = 1
		} else {
			s2Map[string(v2)] = s2Map[string(v2)] + 1
		}
	}

	for key, _ := range s1Map {
		_, exists := s2Map[key]
		if exists {
			return "YES"
		}
	}

	return "NO"
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
