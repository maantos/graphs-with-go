package main

import (
	"fmt"
	"math"
)

type graph struct {
	adjList     map[int][]int
	vertexColor map[int]int
	prevVertex  map[int]int
	sourceDist  map[int]int
}

func CreateGraph() *graph {
	return &graph{
		adjList:     make(map[int][]int),
		vertexColor: make(map[int]int),
		prevVertex:  make(map[int]int),
		sourceDist:  make(map[int]int),
	}
}

func (g *graph) BFS(source int) {
	for i := 0; i < len(g.adjList); i++ {
		if i == source {
			continue
		}
		g.prevVertex[i] = math.MinInt
		g.vertexColor[i] = 0
		g.sourceDist[i] = math.MaxInt
	}
	g.prevVertex[source] = math.MinInt
	g.vertexColor[source] = 1
	g.sourceDist[source] = 0

	queue := make(chan int, len(g.adjList))
	queue <- source

	for len(queue) != 0 {
		x := <-queue
		for _, v := range g.adjList[x] {
			if g.vertexColor[v] == 0 {
				g.vertexColor[v] = 1
				g.sourceDist[v] = g.sourceDist[x] + 1
				g.prevVertex[v] = x
				queue <- v
			}
		}
		g.vertexColor[x] = 1
	}
	defer close(queue)
}

func (g *graph) printPath(start int, end int) {
	if len(g.prevVertex) == 0 {
		fmt.Print("Graph not initialized.")
		return
	}
	if start == end {
		fmt.Print(start)
	} else if g.prevVertex[end] == math.MinInt {
		fmt.Println("No path exists.")
		return
	} else {
		g.printPath(start, g.prevVertex[end])
		fmt.Print(end)
	}
}
