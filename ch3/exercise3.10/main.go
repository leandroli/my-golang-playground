package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("123456789"))
}

func comma(s string) string {
	buf := bytes.Buffer{}
	n := len(s)
	if n <= 3 {
		return s
	}
	buf.WriteString(s[:len(s)%3])
	if len(s)%3 != 0 {
		buf.WriteByte(',')
	}
	for i := len(s) % 3; i < len(s)-3; i += 3 {
		buf.WriteString(s[i : i+3])
		buf.WriteByte(',')
	}
	buf.WriteString(s[len(s)-3:])
	return buf.String()
}
