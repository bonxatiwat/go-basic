package main

import "fmt"

func main() {
	// b := new(int)  // allocate *
	// c := new(*int) // allocate **

	// a := 1 // assign
	// b = &a // point to a
	// c = &b // point to b

	// fmt.Printf("address a %v\n", &a)
	// fmt.Printf("address b %v\n", &b)
	// fmt.Printf("address c %v\n", &c)

	// fmt.Println()

	// fmt.Printf("value a %v\n", a)
	// fmt.Printf("value b %v\n", b)
	// fmt.Printf("value c %v\n", c)

	// fmt.Println()

	// fmt.Printf("*b %v\n", *b)
	// fmt.Printf("*c %v\n", *c)
	// fmt.Printf("**c %v\n", **c)
	a := 2
	// var n *int = &a
	double(&a)
	// fmt.Println(*n)
	fmt.Println(a)
}

// pass by pointers
func double(n *int) {
	*n *= 2
}
