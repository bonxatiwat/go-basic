package main

import "fmt"

func main() {
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	var sum int
	for i := 1; i <= 10; i++ {
		sum += 1
	}

	// For go v1.22+
	for i := range 3 {
		fmt.Println(i)
	}

}
