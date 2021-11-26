package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//fmt.Printf("dec: %d - bin: %08b\n", pc[i], pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

/*
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		fmt.Print("x ")
	}
	return count
}
*/

/*
func PopCount(x uint64) int {
	var count uint64
	for ; x != 0; x >>= 1 {
		count += x & 1
	}
	return int(count)
}
*/

// 2.5 - x&(x-1) removes the right most 1
/*
func PopCount(x uint64) int {
	var count int
	for ; x > 0; x, count = x&(x-1), count+1 {
	}
	return count
}
*/
