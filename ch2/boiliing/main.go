package main

import "fmt"

const bolingF = 212.0

func main() {
	var f float64 = bolingF
	var c float64 = (f - 32) * 5 / 9

	fmt.Printf("boiling point = %g°F or %g°C", f, c)
}
