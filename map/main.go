package main

import "fmt"

var graph = map[int][]int{
	1: {2, 3},
	2: {4, 5},
	3: {6},
}

func main() {
	// a := map[int]int{
	// 	1: 1,
	// 	2: 2,
	// }

	// fmt.Println(a)

	// if _, ok := a[1]; ok {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("Not ok")
	// }

	// fmt.Println(graph)

	// a := make(map[int]int)

	// a[1] = 1

	dfs(graph, 1, make(map[int]bool))
}

func dfs(graph map[int][]int, vertex int, visted map[int]bool) {
	visted[vertex] = true

	for _, v := range graph[vertex] {
		fmt.Printf("->%d", v)
		dfs(graph, v, visted)
	}

}
