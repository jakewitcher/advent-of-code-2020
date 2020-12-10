package jolt

import (
	g "day-10/internal/graph"
	"sort"
)

func CalculateDifference(adapters []int) int {
	adapters = initializeAdapters(adapters)

	joltDifference := make(map[int]int)
	for i := 0; i < len(adapters)-1; i++ {
		difference := adapters[i+1] - adapters[i]
		joltDifference[difference]++
	}

	return joltDifference[1] * joltDifference[3]
}

func CalculatePaths(adapters []int) int {
	adapters = initializeAdapters(adapters)

	graph := g.Graph{Vertices: make([]g.Vertex, len(adapters))}

	for i, j := 0, 1; j < len(adapters); i, j = i+1, i+2 {
		for j < len(adapters) && adapters[j]-adapters[i] <= 3 {
			graph.AddAdjacent(i, g.AdjacentVertex(j))
			j++
		}
	}

	paths := graph.GetPaths()

	return paths
}

func initializeAdapters(adapters []int) []int {
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	return adapters
}
