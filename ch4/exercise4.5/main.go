package main

import "fmt"

func eliminateDup(strings []string) []string {
	out := strings[:0]
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			continue
		}
		out = append(out, strings[i])
	}
	end := len(strings) - 2
	if strings[end] == strings[end+1] {
		out = append(out, strings[end])
	} else {
		out = append(out, strings[end+1])
	}
	return out
}

func main() {
	strings := []string{"dup", "dup", "duplicate", "duplicate"}
	fmt.Println(strings)
	fmt.Println(eliminateDup(strings))
	fmt.Println(strings) // in-place
}
