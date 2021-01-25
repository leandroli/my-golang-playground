package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WordCounter int
type LineCounter int

func scan(p []byte, fn bufio.SplitFunc) (count int) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(fn)
	count = 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}
	return
}

func (c *WordCounter) Write(p []byte) (int, error) {
	count := scan(p, bufio.ScanWords)
	*c += WordCounter(count)
	return count, nil
}

func (c *LineCounter) Write(p []byte) (int, error) {
	count := scan(p, bufio.ScanLines)
	*c += LineCounter(count)
	return count, nil
}

func main() {
	var c WordCounter
	var l LineCounter
	c.Write([]byte("hello"))
	fmt.Println(c)
	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println("Word counter: ", c)

	fmt.Fprintf(&l, "Hello %s\nThis \n is \na line\n.\n.\n", name)
	fmt.Println("Line counter: ", l)
}
