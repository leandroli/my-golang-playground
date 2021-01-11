package main

import "fmt"

func main() {
	fmt.Printf("main: Pointer that f returned: %x\n", f())
	fmt.Println(f() == f())
}

//create a local variable and return it's address
func f() *int {
	var i int
	fmt.Printf("f: Address of local variable: %x\n", &i)
	return &i
}
