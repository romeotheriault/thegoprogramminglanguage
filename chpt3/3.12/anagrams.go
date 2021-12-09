package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need to pass in two arguments.")
		os.Exit(0)
	}
	fmt.Printf("  %t\n", anagram(os.Args[1], os.Args[2]))
}

func anagram(a, b string) bool {
	if a == b {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	m := make(map[rune]int)
	for _, char := range a {
		m[char]++
	}
	for _, char := range b {
		if m[char] == 0 {
			return false
		}
	}
	return true
}
