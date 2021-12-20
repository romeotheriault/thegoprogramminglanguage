package main

// Charcount computes counts of Unicode characters.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	letters := make(map[rune]int) // counts of Unicode characters
	numbers := make(map[rune]int)
	space := make(map[rune]int)
	digits := make(map[rune]int)
	graphic := make(map[rune]int)
	punctuation := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters[r]++
		} else if unicode.IsNumber(r) {
			numbers[r]++
		} else if unicode.IsSpace(r) {
			space[r]++
		} else if unicode.IsDigit(r) {
			digits[r]++
		} else if unicode.IsPunct(r) {
			punctuation[r]++
		} else if unicode.IsGraphic(r) {
			graphic[r]++
		}
		utflen[n]++
	}
	fmt.Printf("Letter\tcount\n")
	for c, n := range letters {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Number\tcount\n")
	for c, n := range numbers {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Space\tcount\n")
	for c, n := range space {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Digit\tcount\n")
	for c, n := range digits {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Graphic\tcount\n")
	for c, n := range graphic {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("Punct\tcount\n")
	for c, n := range punctuation {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
