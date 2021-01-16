package main

import (
	"fmt"
	"gopl.io/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues\n", result.TotalCount)
	fmt.Printf("less than a month before:\n")
	for _, item := range result.Items {
		if time.Now().Before(item.CreatedAt.AddDate(0, 1, 0)) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Printf("less than a year before:\n")
	for _, item := range result.Items {
		if time.Now().Before(item.CreatedAt.AddDate(1, 0, 1)) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Printf("more than a year before:\n")
	for _, item := range result.Items {
		if time.Now().After(item.CreatedAt.AddDate(1, 0, 0)) {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
