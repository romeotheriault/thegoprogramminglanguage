package main

import "fmt"

// Rewrite reverse to use an array pointer instead of a slice.

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%x\n", reverse(&x))
}

func reverse(x *[5]int) *[5]int {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	return x
}
