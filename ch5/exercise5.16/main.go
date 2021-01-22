package main

import "fmt"

func main() {
	words := []string{"May", "you", "have", "a", "strong", "foundation", "when", "the", "wind", "of", "change", "blow"}
	fmt.Println(joinVariadic(" ", words...))
}

func joinVariadic(sep string, strings ...string) string {
	result := ""
	for _, str := range strings {
		result += sep + str
	}
	return result[len(sep):]
}
