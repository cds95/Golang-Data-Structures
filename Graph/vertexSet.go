package graph

type VertexSet struct {
  vertices map[Vertex]bool
}

func (v *VertexSet) NewVertexSet() {
  v.vertices = make(map[Vertex]bool)
}

func (v *VertexSet) Add(vertex Vertex) bool {
  if v.Contains(vertex) {
    return false
  }
  v.vertices[vertex] = true
  return true
}

func (v *VertexSet) Remove(vertex Vertex) {
  delete(v.vertices, vertex)
}

func (v *VertexSet) Contains(vertex Vertex) bool {
  return v.vertices[vertex] == true
}
