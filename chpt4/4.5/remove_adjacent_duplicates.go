package main

import "fmt"

func main() {
	s := []string{"joey", "bobby", "romeo", "romeo", "romeo", "maya", "maya"}
	s = removeAdjacentDuplicates(s)
	fmt.Println(s)
}

func removeAdjacentDuplicates(s []string) []string {
	s2 := s[:0]
	var prev string
	for _, i := range s {
		if i != prev {
			s2 = append(s2, i)
		}
		prev = i
	}
	return s2
}
