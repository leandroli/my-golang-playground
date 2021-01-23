package main

import (
	"bytes"
	"fmt"
)

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(1&i)
	}
}

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

//!+len

// Len return the number of elements
func (s *IntSet) Len() int {
	sumNum := 0
	for word := range s.words {
		for i := 0; i < 8; i++ {
			sumNum += int(pc[byte(word>>(i*8))])
		}
	}
	return sumNum
}

//!-len

//!+remove

//Remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

//!-remove

//!+clear

//Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

//!-clear

//!+copy

//Copy return a copy for the set
func (s *IntSet) Copy() *IntSet {
	result := make([]uint64, len(s.words))
	copy(result, s.words)
	return &IntSet{words: result}
}

//!-copy

func main() {
	var set IntSet
	set.Add(1)
	set.Add(155)
	set.Add(177)
	fmt.Println(set.String())
	setCopy := set.Copy()
	setCopy.Add(178)
	fmt.Println(set.String())
	fmt.Println(setCopy)
	set.Clear()
	fmt.Println(set)
	fmt.Println(setCopy)
}
