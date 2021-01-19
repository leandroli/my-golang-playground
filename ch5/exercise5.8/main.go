package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) (bool, *html.Node) {
	if pre != nil {
		if !pre(n) {
			return false, n
		}
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		if ok, node := forEachNode(i, pre, post); !ok {
			return false, node
		}
	}
	if post != nil {
		if !post(n) {
			return false, n
		}
	}
	return true, nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	_, elem := forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, item := range n.Attr {
				if item.Key == "id" && item.Val == id {
					return false
				}
			}
		}
		return true
	}, func(n *html.Node) bool {
		return true
	})
	return elem
}

func main() {
	url := "https://www.bupt.edu.cn"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise5.8: %v", err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "exercise5.8: getting %s: %s", url, resp.Status)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise5.8: parsing %s as HTML: %v", url, err)
		os.Exit(1)
	}
	fmt.Println(ElementByID(doc, "section-intro"))
}
