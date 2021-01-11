package main

import (
	"container/list"
	"fmt"
)

var Visited []bool

// Initialize the graph by entering number of nodes
func buildGraph(v int) *list.List {
	MasterList := list.New()
	for i := 0; i < v; i++ {
		Visited = make([]bool, v)
		var l = list.New()
		l.PushBack(i)
		MasterList.PushBack(l)
	}
	return MasterList
}

// Create an edge of the graph
func addEdge(src int, dst int, graph *list.List) {

	for v := graph.Front(); v != nil; v = v.Next() {
		if v.Value.(*list.List).Front().Value == src {
			v.Value.(*list.List).PushBack(dst)
		}
	}

}

// Build a custom FIFO queue
type Queue struct {
	values []interface{}
}

func New() *Queue {
	return &Queue{
		values: make([]interface{}, 0),
	}
}

// Push element at the back of queue
func (q *Queue) Push(key interface{}) {
	q.values = append(q.values, key)
}

// Pop the first element of array
func (q *Queue) Pop() interface{} {
	if len(q.values) == 0 {
		return nil
	}
	returnValue := q.values[0]
	q.values = q.values[1:]
	return returnValue
}

func (q *Queue) Length() int {
	return len(q.values)
}

// Perform Breadth first search
func Bfs(src int, graph *list.List) {
	q := New()
	q.Push(src)
	Visited[src] = true
	//fmt.Println(q.Length())
	//fmt.Println(visited)

	for q.Length() != 0 {
		elem := q.Pop()
		fmt.Println(elem)
		for e := graph.Front(); e != nil; e = e.Next() { // Find the adjacent nodes
			if e.Value.(*list.List).Front().Value == elem {

				for qv := e.Value.(*list.List).Front(); qv != nil; qv = qv.Next() { // Push all adjacent nodes to queue
					if Visited[qv.Value.(int)] != true {
						q.Push(qv.Value)
						Visited[qv.Value.(int)] = true
					}
				}
				break
			}
		}
	}

}

func main() {
	fmt.Println("Breadth first search")
	var Output = buildGraph(4)
	//fmt.Println(visited)
	//fmt.Println(Output.Len())
	addEdge(0, 1, Output)
	addEdge(0, 2, Output)
	addEdge(1, 2, Output)
	addEdge(2, 0, Output)
	addEdge(2, 3, Output)
	addEdge(3, 3, Output)

	fmt.Println("Start BFS")

	Bfs(2, Output)

}
