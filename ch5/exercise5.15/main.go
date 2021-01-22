package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 7, 2}
	fmt.Println(max(values...))

}

func max(vals ...int) int {
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min(vals ...int) int {
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max2(m int, vals ...int) int {
	max := m
	for _, val := range vals {
		if max < val {
			max = val
		}
	}
	return max
}

func min2(m int, vals ...int) int {
	min := m
	for _, val := range vals {
		if min > val {
			min = val
		}
	}
	return min
}
