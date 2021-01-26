package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type stringReader struct {
	s string
}

func (r *stringReader) Read(p []byte) (n int, err error) {
	n = copy(p, r.s)
	r.s = r.s[n:]
	if len(r.s) == 0 {
		err = io.EOF
	}
	return
}

func NewReader(s string) (newReader *stringReader) {
	newReader = &stringReader{s: s}
	return
}

func main() {
	s := "<html><body><p>hi</p></body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		fmt.Println("exercise7.4: ", err)
	}
}
