package jolt

type Edge int
type Vertex []Edge

type Graph struct {
	vertices []Vertex
}

func (g *Graph) AddEdge(v int, e Edge) {
	g.vertices[v] = append(g.vertices[v], e)
}

func (g *Graph) GetPaths() int {
	seen := &SeenPaths{edges: make(map[Edge]int)}
	return g.getPaths(g.vertices[0], seen)
}

func (g *Graph) getPaths(vertex Vertex, seen *SeenPaths) int {
	if len(vertex) == 0 {
		return 1
	}

	var vertexPaths int
	for _, edge := range vertex {
		if count, ok := seen.HasKey(edge); ok {
			vertexPaths += count
			continue
		}

		edgePaths := g.getPaths(g.vertices[edge], seen)
		seen.Add(edge, edgePaths)

		vertexPaths += edgePaths
	}

	return vertexPaths
}

type SeenPaths struct {
	edges map[Edge]int
}

func (s *SeenPaths) Add(edge Edge, paths int) {
	s.edges[edge] = paths
}

func (s *SeenPaths) HasKey(edge Edge) (int, bool) {
	count, ok := s.edges[edge]
	return count, ok
}
