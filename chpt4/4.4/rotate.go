package main

import "fmt"

// Write a version of rotate that operates on a single pass.

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 2)
	fmt.Printf("%x\n", s)
	// Rotate s left by two positions
	// Will be: {2, 3, 4, 5, 0, 1}

	/* Multi-pass version
	reverse(s[:2])   // 1,0,2,3,4,5
	reverse(s[2:])   // 1,0,5,4,3,2
	reverse(s)       // 2,3,4,5,0,1
	*/

}

func rotate(s []int, rotate int) []int {
	if rotate > len(s) {
		rotate = rotate % len(s)
	}
	return append(s[rotate:], s[:rotate]...)
}
