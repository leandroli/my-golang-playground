package main

import "fmt"

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] != 0 {
			return num
		}
		m[num]++
	}
	panic(fmt.Errorf(""))
}
