package main

import "fmt"

func main() {
	s := [32]int{5: 1, 12: 1, 3, 5, 7}
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}

func reverse(ptr *[32]int) {
	for i, j := 0, len(*ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}

}
