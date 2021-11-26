package main

import (
	"fmt"

	p "2.3/popcount"
)

func main() {
	var x uint64 = 2987888811122299999
	fmt.Println(p.PopCount(x))
}
