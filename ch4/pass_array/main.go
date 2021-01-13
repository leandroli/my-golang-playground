package main

import "fmt"

func main() {
	a := [32]byte{5: 1}
	fmt.Println(a)
	zero(&a)
	fmt.Println(a)
}

func zero(ptr *[32]byte) {
	fmt.Println(*ptr)
	fmt.Println(ptr)
	fmt.Println(ptr[5]) // is ptr[5] == (*ptr)[5]? u
	fmt.Println((*ptr)[5])
	for i := range ptr {
		ptr[i] = 0
	}
}
