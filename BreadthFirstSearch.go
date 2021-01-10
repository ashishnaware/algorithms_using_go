package main

import (
	"container/list"
	"fmt"
)

func buildGraph(v int) *list.List {
	MasterList := list.New()
	for i := 0; i < v ; i++{
		var l = list.New()
		l.PushBack(i)
		MasterList.PushBack(l)
	}
	return MasterList
}

func addEdge(src int,dst int, graph *list.List)  {
	
	for v:=graph.Front(); v!= nil; v=v.Next(){
		if v.Value.(*list.List).Front().Value == src {
			v.Value.(*list.List).PushBack(dst)
		}
	}
	
}

func Bfs(src int)  {

}




func main() {
	fmt.Println("Breadth first search")
	var Output = buildGraph(4)
	fmt.Println(Output.Len())
	addEdge(0,1,Output)
	addEdge(0,2,Output)
	addEdge(1,2,Output)
	addEdge(2,0,Output)
	addEdge(2,3,Output)
	addEdge(3,3,Output)

	fmt.Println("Start BFS")


	for e := Output.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*list.List).Len())


	}


}
