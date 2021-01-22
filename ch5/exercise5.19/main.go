package main

import "fmt"

func main() {
	fmt.Println(f())
}

//return a non-zero value without return statement
func f() (i int) {
	defer func() {
		recover()
		i = 1
	}()
	panic("")
}
