package main

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc [256]byte
var loadPcOnce sync.Once

func loadPc() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(Pc(byte(x>>(0*8))) +
		Pc(byte(x>>(1*8))) +
		Pc(byte(x>>(2*8))) +
		Pc(byte(x>>(3*8))) +
		Pc(byte(x>>(4*8))) +
		Pc(byte(x>>(5*8))) +
		Pc(byte(x>>(6*8))) +
		Pc(byte(x>>(7*8))))
}

func Pc(i byte) byte {
	loadPcOnce.Do(loadPc)
	return pc[i]
}

func main() {
	for i := 0; i < 64; i++ {
		fmt.Println(PopCount(uint64(i)))
	}
}
