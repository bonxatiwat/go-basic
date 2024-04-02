package main

import "fmt"

func main() {
	hello("Bonx")
	hello("Ninew")

	result := add(2, 4)

	fmt.Println(result)
	fmt.Println(add(2, 2))
	fmt.Println(sub(2, 1))
	fmt.Println(mul(2, 3))
	fmt.Println(div(8, 0))

	// name := "bonxatiwat"

	result2 := func(a, b int) int {
		// name = "ninew"
		// fmt.Println("Hello, ", name)
		return a + b
	}(1, 2)

	fmt.Println(result2)
}

func hello(name string) {
	fmt.Println("Hello, ", name)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero.")
		return 0
	}
	return a / b
}
