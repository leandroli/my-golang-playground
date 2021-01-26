package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r        io.Reader
	limit, n int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	n, err = lr.r.Read(p[:lr.limit-lr.n])
	lr.n += int64(n)
	if lr.n >= lr.limit {
		err = io.EOF
	}
	return
}

func LimitReader(r io.Reader, n int64) (lr io.Reader) {
	lr = &limitReader{r: r, limit: n}
	return
}

func main() {
	s := "hi there"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 4)
	n, err := b.ReadFrom(r)
	fmt.Println("exercise7.5:", "n =", n, "err =", err, "b.String() =", b.String())
}
