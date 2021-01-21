package main

import (
	"fmt"
	"os"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
	//"linear algebra": {"calculus"}, // additional for this exercise
}

//!-table

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise5.11: %v", err)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	seen := make(map[string]bool)
	checkCircle := make(map[string]bool) // check whether there is a circle for each path
	var visitAll func([]string) error
	var sequence []string
	visitAll = func(items []string) error {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				checkCircle[item] = true
				err := visitAll(m[item])
				if err != nil {
					return err
				}
				checkCircle[item] = false
				sequence = append(sequence, item)
			} else {
				if checkCircle[item] {
					return fmt.Errorf("this directed graph has circle(s)")
				}
			}
		}
		return nil
	}

	var items []string
	for i := range m {
		items = append(items, i)
	}
	err := visitAll(items)
	if err != nil {
		return nil, fmt.Errorf("topoSort: %v", err)
	}
	return sequence, nil
}
