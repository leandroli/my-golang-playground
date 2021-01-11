package main

import "fmt"

func popCount(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if byte((x>>i)&((x>>i)-1)<<7) != byte((x>>i)<<7) { // shifting left by 7 bits is to get rightmost bit
			count++
		}
	}
	return count
}

func main() {
	for i := 0; i < 64; i++ {
		fmt.Println(popCount(uint64(i)))
	}
}
