package main

import (
	"fmt"
	"sync"
)

func main() {
	total := 0
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			total += i
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("total:", total)
}
