package main

import (
	"bufio"
	"fmt"
	"os"
)

type WordPair struct {
	word  string
	count int
}

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

	var wp []WordPair

	for k, v := range wordfreq {
		//fmt.Printf("%20s: %d\n", k, v)
		var j WordPair
		j.count = v
		j.word = k
		wp = append(wp, j)
	}

	// Quicksort the wp (word pair) array based on value (word count)
	qsort(wp, 0, len(wp)-1)
	for _, v := range wp {
		fmt.Printf("%20s %d\n", v.word, v.count)
	}
}

func qsort(wp []WordPair, l int, r int) {
	if r-l <= 0 {
		return
	}
	var pivot int = qpartition(wp, l, r)
	qsort(wp, l, pivot-1)
	qsort(wp, pivot+1, r)
}

func qpartition(wp []WordPair, l int, r int) int {
	var pivot int = r
	r = r - 1

	for {
		for wp[l].count < wp[pivot].count {
			l++
		}
		for wp[r].count > wp[pivot].count && r > 0 {
			r--
		}
		if l >= r {
			break
		}
		// Swap left and right values
		wp[l], wp[r] = wp[r], wp[l]
		l++
	}

	// Swap the pivot with the left pointer (it's final destination)
	wp[pivot], wp[l] = wp[l], wp[pivot]
	return l
}
