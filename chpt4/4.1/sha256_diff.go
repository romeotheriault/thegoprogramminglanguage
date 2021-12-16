package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	sha256BitDiff(&c1, &c2)
	// To get the number of different bits you need to:
	// XOR (^) the two sha256 bytes together,
	// Then take the population count (bits.OnesCount8) of the result.
}

func sha256BitDiff(c1 *[32]byte, c2 *[32]byte) {
	var sum int = 0
	for i := range c1 {
		xor := c1[i] ^ c2[i]
		sum += int(pc[byte(xor>>(0*8))])
		//sum += bits.OnesCount8(xor)
	}
	fmt.Printf("Different bits in sha256 sums: %d\n", sum)
}
