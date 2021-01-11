package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omitting trailing newline.")
var sep = flag.String("sep", " ", "separator")

func main() {
	flag.Parse()
	fmt.Printf(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
