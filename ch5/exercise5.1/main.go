package main

import "golang.org/x/net/html"

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for i := n.NextSibling; i != nil; i = n.NextSibling {
		links = visit(links, i)
	}
	return visit(links, n.FirstChild)
}
