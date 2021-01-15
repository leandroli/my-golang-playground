package main

import (
	"bufio"
	"fmt"
	"os"
)

var fileName string = "ch4/exercise4.9/test.txt"

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise4.9: %v", err)
		os.Exit(1)
	}
	counts := map[string]int{}
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Printf("word\tfrequency\n")
	for w, c := range counts {
		fmt.Printf("%s\t%d\n", w, c)
	}
}
