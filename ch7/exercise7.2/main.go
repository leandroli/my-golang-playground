package main

import "io"

type newWriter struct {
	w       io.Writer
	counter int64
}

func (c *newWriter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.counter += int64(n)
	return
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := &newWriter{w, 0}
	return c, &c.counter
}
