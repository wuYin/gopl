// graph 建立一个从字符串到字符串集合的映射
package main

import "fmt"

var graph = make(map[string]map[string]bool)

func main() {
	addEdge("a", "b")
	fmt.Println(hasEdge("a", "c"))
	fmt.Println(hasEdge("a", "b"))
}

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true // 修改引用指向的数据，等同于直接修改 graph[from][to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}
