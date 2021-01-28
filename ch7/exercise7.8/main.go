package exercise7_8

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type columnCmp func(a, b *Person) comparison

type PeopleSort struct {
	p          []Person
	columns    []columnCmp
	maxColumns int
}

func NewPeopleSort(p []Person, maxColumns int) *PeopleSort {
	return &PeopleSort{p, nil, maxColumns}
}

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

type OrderOption int

const (
	ByName OrderOption = iota
	ByAge
	BySumOfAgeDigits
)

func (c *PeopleSort) lessName(a, b *Person) comparison {
	switch {
	case a.Name == b.Name:
		return eq
	case a.Name < b.Name:
		return lt
	default:
		return gt
	}
}

func (c *PeopleSort) lessSumOfAgeDigits(a, b *Person) comparison {
	aSum := sumOfDigits(a.Age)
	bSum := sumOfDigits(b.Age)
	switch {
	case aSum == bSum:
		return eq
	case aSum < bSum:
		return lt
	default:
		return gt
	}
}

func sumOfDigits(n int) int {
	sum := 0
	for ; n > 0; n /= 10 {
		sum += n % 10
	}
	return sum
}

func (c *PeopleSort) lessAge(a, b *Person) comparison {
	switch {
	case a.Age == b.Age:
		return eq
	case a.Age < b.Age:
		return lt
	default:
		return gt
	}
}

func (c *PeopleSort) Len() int      { return len(c.p) }
func (c *PeopleSort) Swap(i, j int) { c.p[i], c.p[j] = c.p[j], c.p[i] }

func (c *PeopleSort) Less(i, j int) bool {
	for _, f := range c.columns {
		cmp := f(&c.p[i], &c.p[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (c *PeopleSort) Select(orderOption OrderOption) {
	// Prepend the new comparison, as it's the most significant.
	switch orderOption {
	case ByName:
		c.columns = append([]columnCmp{c.lessName}, c.columns...)
	case ByAge:
		c.columns = append([]columnCmp{c.lessAge}, c.columns...)
	case BySumOfAgeDigits:
		c.columns = append([]columnCmp{c.lessSumOfAgeDigits}, c.columns...)
	}

	// Don't let the slice of comparisons grow without bound.
	if len(c.columns) > c.maxColumns {
		c.columns = c.columns[:c.maxColumns]
	}
}
