package main

//	1 -- 2   5 -- 6
//	|	|  / |  / |
//	0	3 -- 4 -- 7   8

func main() {
	g := CreateGraph()
	g.adjList[0] = []int{1}
	g.adjList[1] = []int{0, 2}
	g.adjList[2] = []int{1, 3}
	g.adjList[3] = []int{2, 4, 5}
	g.adjList[4] = []int{3, 5, 7}
	g.adjList[5] = []int{3, 4, 6}
	g.adjList[6] = []int{4, 5, 7}
	g.adjList[7] = []int{4, 6}
	g.adjList[8] = []int{}

	g.BFS(8)

	g.printPath(3, 8)

}
