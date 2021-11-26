package main

import (
	p "2.3/popcount"
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		p.PopCount(2987)
	}
}
