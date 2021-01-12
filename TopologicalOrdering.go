package main

import (
	"container/list"
	"fmt"
)

var degree []int

// Initialize the graph by entering number of nodes
func constructGraph(v int) *list.List {
	degree = make([]int,v)
	MasterList := list.New()
	for i := 0; i < v; i++ {
		var l = list.New()
		l.PushBack(i)
		MasterList.PushBack(l)
	}
	return MasterList
}

// Create an edge of the graph
func addTEdge(src int, dst int, graph *list.List) {

	for v := graph.Front(); v != nil; v = v.Next() {
		if v.Value.(*list.List).Front().Value == src {
			v.Value.(*list.List).PushBack(dst)
		}
	}

}

func TopologicalOrdering(graph *list.List)  {

	for e:= graph.Front(); e != nil; e = e.Next(){
		front := e.Value.(*list.List).Front()
		for e1 := front.Next(); e1 != nil; e1 = e1.Next(){
			degree[e1.Value.(int)] += 1 	// Increase the degree of vertex by 1
		}
	}
	fmt.Println(degree)
}


func main() {
	fmt.Println("Topological ordering")
	var Output = constructGraph(6)

	addTEdge(5, 2, Output)
	addTEdge(5, 0, Output)
	addTEdge(4, 0, Output)
	addTEdge(4, 1, Output)
	addTEdge(2, 3, Output)
	addTEdge(3, 1, Output)
	TopologicalOrdering(Output)
	//fmt.Println("Start DFS")

}
