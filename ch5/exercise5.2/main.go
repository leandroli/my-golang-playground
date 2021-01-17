package main

import "golang.org/x/net/html"

func elemNameNum(elemNums map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		elemNums[n.Data]++
	}
	for i := n.FirstChild; i != nil; i = i.NextSibling {
		elemNums = elemNameNum(elemNums, i)
	}
	return elemNums
}
