package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	timeZonesAndServerAddr := make(map[string]string)
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Usage: clockwall NAME=ADDR")
	}
	for _, arg := range flag.Args() {
		temp := strings.Split(arg, "=")
		if len(temp) != 2 {
			fmt.Fprintf(os.Stderr, "Bad argument, pass arguemt as NAME=ADDR")
		}
		timeZonesAndServerAddr[temp[0]] = timeZonesAndServerAddr[temp[1]]
	}
	for timeZone, addr := range timeZonesAndServerAddr {
		go handleConn(timeZone, addr)
	}
	for {
	}
}

func handleConn(timeZone, addr string) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	buf := new(bytes.Buffer)
	mustCopy(buf, c)
	fmt.Println(timeZone, buf.String())
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
