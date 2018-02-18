package graph

type EdgeSet struct {
  edges map[Edge]bool
}

func (e *EdgeSet) Add(edge Edge) {
  e.edges[edge] = true
}

func (e *EdgeSet) Remove(edge Edge) {
  delete(e.edges, edge)
}

func (e *EdgeSet) Contains(edge Edge) bool {
  return e.edges[edge] == true
}

func (e *EdgeSet) NewEdgeSet() {
  e.edges = make(map[Edge]bool)
}
