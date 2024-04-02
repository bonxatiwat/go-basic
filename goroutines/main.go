package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := range 3 {
		go func(index int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(index)
		}(i, &wg)
	}

	wg.Wait()
}
