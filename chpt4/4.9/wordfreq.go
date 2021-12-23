package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("os.Open: ", err)
		os.Exit(1)
	}

	var wordfreq = make(map[string]int)

	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		wordfreq[input.Text()]++
	}

	for k, v := range wordfreq {
		fmt.Printf("%20s: %d\n", k, v)
	}
}
