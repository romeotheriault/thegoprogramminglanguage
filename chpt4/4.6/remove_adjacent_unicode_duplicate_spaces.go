package main

import (
	"fmt"
	"unicode"
)

/*
ASCII space, U+0020

Squash adjacent unicode space's (unicode.IsSpace) in a UTF-8 encoded []byte slice into a single ascii space (U+0020)
*/

func main() {
	s := "Wow     look     at me   !   "
	bs := []byte(s)
	bs = removeAdjacentDuplicateUnicodeSpaces(bs)
	fmt.Printf("%s\n", bs)
}

func removeAdjacentDuplicateUnicodeSpaces(bs []byte) []byte {
	bs2 := bs[:0]
	prev := false
	for _, b := range bs {
		if !unicode.IsSpace(rune(b)) {
			bs2 = append(bs2, b)
			prev = false
		} else {
			if prev == false {
				bs2 = append(bs2, 32)
			}
			prev = true
		}
	}
	return bs2
}
