package main

import (
	"fmt"

	p "2.3/popcount"
)

func main() {
	var x uint64 = 22
	fmt.Println(p.PopCount(x))
	/*
		var x uint64 = 22
		var count int
		for ; x > 0; x, count = x&(x-1), count+1 {
		}
	*/

	//fmt.Println(count)

	//fmt.Printf("%08b\n", 22)
	//fmt.Printf("%08b\n", 22&(22-1))
}
