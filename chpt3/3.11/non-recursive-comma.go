package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

/*
Update comma so it deals with floating point numbers and an optional sign.

Write a non-recursive version of comma, using bytes.Buffer instead of
string concatenation.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
*/

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	slen := len(s) - 1

	// If we're less than 3 we know there are no commas
	if slen < 3 {
		return s
	}

	// Check if there's a sign.
	var sign bool = false
	if s[0] == '-' || s[0] == '+' {
		sign = true
		//fmt.Println(sign)
	}

	// Check if we're a float.
	idx := strings.Index(s, ".")
	if idx != -1 {
		for len := slen; len >= idx; len-- {
			buf.WriteByte(s[len])
		}
		slen = idx - 1
	}

	min := 0
	if sign {
		min = 1
	}
	for i, count := slen, 1; i >= min; i, count = i-1, count+1 {
		if count%3 == 0 && i != min {
			buf.WriteByte(s[i])
			buf.WriteByte(',')
		} else {
			buf.WriteByte(s[i])
		}
	}

	// Reverse the byte array.
	sc := buf.Bytes()
	for i, j := 0, len(sc)-1; i < j; i, j = i+1, j-1 {
		sc[i], sc[j] = sc[j], sc[i]
	}
	// If we have a sign, prepend it, otherwise just return the string.
	if sign {
		return string(s[0]) + string(sc)
	} else {
		return string(sc)
	}
}
