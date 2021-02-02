package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	depth := flag.Int("depth", 3, "URLs that are reachable by at most three links will be fetched.")
	type item struct {
		depth int
		links []string
	}
	workList := make(chan item)
	//var wg sync.WaitGroup
	var n int
	flag.Parse()
	n++
	go func() { workList <- item{depth: 0, links: flag.Args()} }()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()
	//go func() {
	//	wg.Wait()
	//	close(workList)
	//}()
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		msg := <-workList
		if msg.depth > *depth {
			continue
		}
		for _, link := range msg.links {
			if !seen[link] {
				seen[link] = true
				n++
				//wg.Add(1)
				go func(msg item, link string) {
					//defer wg.Done()
					workList <- item{depth: msg.depth + 1, links: crawl(link)}
				}(msg, link)
			}
		}
	}
}

var tokens = make(chan struct{}, 20)

func crawl(link string) []string {
	log.Println(link)
	select {
	case <-cancel:
		return nil
	case tokens <- struct{}{}:
		list, err := Extract(link)
		<-tokens
		if err != nil {
			log.Print(err)
		}
		return list
	}
}

var cancel = make(chan struct{})

// Extract makes an HTTP GET request to the specified URL, parses
// the response as HTML, and returns the links in the HTML document.
func Extract(url string) ([]string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
