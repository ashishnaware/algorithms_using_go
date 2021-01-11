package main

import (
	"container/list"
	"fmt"
)

var dfsVisited []bool

// Initialize the graph by entering number of nodes
func buildDFSGraph(v int) *list.List {
	MasterList := list.New()
	for i := 0; i < v; i++ {
		dfsVisited = make([]bool, v)
		var l = list.New()
		l.PushBack(i)
		MasterList.PushBack(l)
	}
	return MasterList
}

// Create an edge of the graph
func addDFSEdge(src int, dst int, graph *list.List) {

	for v := graph.Front(); v != nil; v = v.Next() {
		if v.Value.(*list.List).Front().Value == src {
			v.Value.(*list.List).PushBack(dst)
		}
	}

}

func recursiveDFS(s int,graph *list.List) {
	for e:= graph.Front();e != nil; e = e.Next(){
		if e.Value.(*list.List).Front().Value == s {
			for e1:= e.Value.(*list.List).Front(); e1 != nil; e1 = e1.Next(){
				if dfsVisited[e1.Value.(int)] != true{
					fmt.Println(e1.Value.(int))
					dfsVisited[e1.Value.(int)] = true
					recursiveDFS(e1.Value.(int),graph)
				}
			}

		}
	}
}

func DFS(s int,graph *list.List) {
	dfsVisited[s] = true
	fmt.Println(s)
	recursiveDFS(s,graph)
}

func main() {
	fmt.Println("Depth first search")
	var Output = buildDFSGraph(4)
	//fmt.Println(visited)
	//fmt.Println(Output.Len())
	addDFSEdge(0, 1, Output)
	addDFSEdge(0, 2, Output)
	addDFSEdge(1, 2, Output)
	addDFSEdge(2, 0, Output)
	addDFSEdge(2, 3, Output)
	addDFSEdge(3, 3, Output)

	fmt.Println("Start DFS")
	DFS(0,Output)
	fmt.Println(Output)
	//Bfs(2, Output)

}
