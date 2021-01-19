package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func myAtoi(s string) int {
	nonspace := s
	negative := false
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			nonspace = s[i+1:]
		} else {
			break
		}
	}
	if nonspace == "" || (!unicode.IsDigit(rune(nonspace[0])) && nonspace[0] != '+' && nonspace[0] != '-') {
		return 0
	}
	if nonspace[0] == '-' {
		negative = true
		nonspace = nonspace[1:]
	} else if nonspace[0] == '+' {
		nonspace = nonspace[1:]
	}
	for i := 0; i < len(nonspace); i++ {
		if !unicode.IsDigit(rune(nonspace[i])) {
			nonspace = nonspace[0:i]
		}
	}
	result, _ := strconv.Atoi(nonspace)
	if negative {
		return -result
	}
	return result
}

func main() {
	fmt.Println(myAtoi("4193 with words"))
}
