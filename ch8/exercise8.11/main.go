package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- request("stackoverflow.com") }()
	go func() { responses <- request("www.baidu.com") }()
	go func() { responses <- request("github.com") }()
	firstResp := <-responses // return the quickest response
	close(cancel)
	return firstResp
}

var cancel = make(chan struct{})

func request(hostname string) (response string) {
	req, err := http.NewRequest("Get", "https://"+hostname, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("getting %s: %s", hostname, resp.Status)
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		log.Println(err)
		return ""
	}
	return buf.String()
}

func main() {
	fmt.Println(mirroredQuery())
}
