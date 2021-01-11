package main

import "fmt"

func exercise2_4(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if ((x >> i) & 1) == 1 {
			count++
		}
	}
	return count
}

func main() {
	for i := 0; i < 1024; i++ {
		fmt.Println(exercise2_4(uint64(i)))
	}
}
