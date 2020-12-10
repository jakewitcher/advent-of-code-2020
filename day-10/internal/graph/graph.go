package graph

type AdjacentVertex int
type Vertex []AdjacentVertex

type Graph struct {
	Vertices []Vertex
}

func (g *Graph) AddAdjacent(v int, av AdjacentVertex) {
	g.Vertices[v] = append(g.Vertices[v], av)
}

func (g *Graph) GetPaths() int {
	seen := &SeenPaths{vertices: make(map[AdjacentVertex]int)}
	return g.getPaths(g.Vertices[0], seen)
}

func (g *Graph) getPaths(vertex Vertex, seen *SeenPaths) int {
	if len(vertex) == 0 {
		return 1
	}

	var vertexPaths int
	for _, adj := range vertex {
		if count, ok := seen.HasKey(adj); ok {
			vertexPaths += count
			continue
		}

		adjPaths := g.getPaths(g.Vertices[adj], seen)
		seen.Add(adj, adjPaths)

		vertexPaths += adjPaths
	}

	return vertexPaths
}

type SeenPaths struct {
	vertices map[AdjacentVertex]int
}

func (s *SeenPaths) Add(av AdjacentVertex, paths int) {
	s.vertices[av] = paths
}

func (s *SeenPaths) HasKey(av AdjacentVertex) (int, bool) {
	count, ok := s.vertices[av]
	return count, ok
}
