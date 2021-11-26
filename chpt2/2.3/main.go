package main

import (
	"fmt"

	p "2.3/popcount"
)

func main() {
	var x uint64 = 22
	fmt.Println(p.PopCount(x))
}
