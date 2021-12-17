package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Printf("%s print SHA hash of stdin input. Sha256 by default, 384 and 512 with flag.\n", os.Args[0])
	fmt.Printf("%s [-384|-512]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	arglen := len(os.Args)
	stdin := bufio.NewScanner(os.Stdin)

	// by default, bufio.Scanner scans newline-separated lines
	var b []byte
	for stdin.Scan() {
		b = stdin.Bytes()
	}
	if arglen == 1 {
		fmt.Printf("%x\n", sha256.Sum256(b))
		os.Exit(0)
	} else if arglen == 2 {
		switch os.Args[1] {
		case "-384":
			fmt.Printf("%x\n", sha512.Sum384(b))
		case "-512":
			fmt.Printf("%x\n", sha512.Sum512(b))
		default:
			usage()
		}
	} else {
		usage()
	}
	/*
		c1 := sha256.Sum256([]byte("x"))
		c2 := sha512.Sum384([]byte("X"))
		c3 := sha512.Sum512([]byte("X"))
		fmt.Printf("c1: %x\n", c1)
		fmt.Printf("c2: %x\n", c2)
		fmt.Printf("c3: %x\n", c3)
	*/
}
