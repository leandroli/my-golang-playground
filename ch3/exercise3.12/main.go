package main

import "fmt"

func main() {
	fmt.Println(isAnagrams("quite", "quiet"))
	fmt.Println(isAnagrams("anagram", "calamitous"))
}

func isAnagrams(s1, s2 string) bool {
	return equalMap(count(s1), count(s2))
}

func count(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, c := range s {
		counts[c]++
	}
	return counts
}

func equalMap(m1, m2 map[rune]int) bool {
	for i, v := range m1 {
		if m2[i] != v {
			return false
		}
	}
	for i, v := range m2 {
		if m1[i] != v {
			return false
		}
	}
	return true
}
