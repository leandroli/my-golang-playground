package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := map[rune]int{}
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	letterCount := 0
	digitCount := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "exercise4.8: %v", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			digitCount++
		} else if unicode.IsLetter(r) {
			letterCount++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, c := range counts {
		fmt.Printf("%q\t%d\n", r, c)
	}
	fmt.Printf("utflen\tcount\n")
	for i, v := range utflen {
		fmt.Printf("%d\t%d\n", i, v)
	}
	fmt.Printf("\n%d digital characters.\n", digitCount)
	fmt.Printf("\n%d letters.\n", letterCount)
	fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
}
