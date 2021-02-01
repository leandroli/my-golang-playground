package main

import (
	"flag"
	"gopl.io/ch5/links"
	"log"
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
	tokens <- struct{}{}
	list, err := links.Extract(link)
	<-tokens
	if err != nil {
		log.Print(err)
	}
	return list
}
