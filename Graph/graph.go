package graph

import (
  "fmt"
)

type Graph struct {
  vertices VertexSet
  adjacencies map[int][]Vertex
  edges EdgeSet
  VerticesCount int
  EdgesCount int
}

// Adds a vertex to the graph and returns true if successful.  Method fails if
// graph already contains the given vertex
func (g *Graph) AddVertex(v Vertex) bool {
  if g.VerticesCount == 0 {
    g.vertices.NewVertexSet()
  }
  g.VerticesCount++
  return g.vertices.Add(v)
}

// Checks to see if graph contains a vertex.  Returns true if it contains it
// and false otherwise
func (g *Graph) ContainsVertex(v Vertex) bool {
  return g.vertices.Contains(v)
}

// Adds an edge to the graph.  Returns true if successful and false otherwise
// Method will fail if edge weight is negative or the graph doesn't contain
// either the source or destination vertex
func (g *Graph) AddEdge(e Edge) bool {
  if !g.ContainsVertex(e.Source) || !g.ContainsVertex(e.Destination) || e.Weight < 0 {
    return false
  }
  if g.EdgesCount == 0 {
    g.edges.NewEdgeSet()
  }
  g.edges.Add(e)
  if g.adjacencies == nil {
    g.adjacencies = make(map[int][]Vertex)
  }
  adjacencies := g.adjacencies[e.Source.Index]
  adjacencies = append(adjacencies, e.Destination)
  g.adjacencies[e.Source.Index] = adjacencies
  return true
}

// Prints out all vertices in graph
func (g *Graph) PrintVertices() {
  for vertex := range g.vertices.vertices {
    fmt.Println(vertex)
  }
}

// Prints all edges in graph
func (g *Graph) PrintEdges() {
  for edge := range g.edges.edges {
    fmt.Println(edge)
  }
}

// Prints out the adjacency lists in the graph
func (g *Graph) PrintAdjacencyList() {
  fmt.Println(g.adjacencies)
}

/*
  Runs the BFS Search algorithm on the graph and prints out the results line
  by line
*/
func (g *Graph) BFS(start Vertex) {
  vertexQueue := []Vertex{start}
  seen := make(map[int]bool)
  seen[start.Index] = true

  for len(vertexQueue) > 0 {
    currentVertex := vertexQueue[0]
    fmt.Println(currentVertex.City)
    vertexQueue = append(vertexQueue[:0], vertexQueue[1:]...)

    for _, vertex := range g.adjacencies[currentVertex.Index] {
      if seen[vertex.Index] == false {
        seen[vertex.Index] = true
        vertexQueue = append(vertexQueue, vertex)
      }
    }
  }
}

/*
  Runs the DFS Search algorithm through the graph and prints out each
  vertex's city
*/
func (g *Graph) DFS(start Vertex) {
  vertexStack := []Vertex{start}
  seen := make(map[int]bool)
  seen[start.Index] = true

  for len(vertexStack) > 0 {
    currentVertex := vertexStack[0]
    fmt.Println(currentVertex.City)
    vertexStack = append(vertexStack[:0], vertexStack[1:]...)

    for _, vertex := range g.adjacencies[currentVertex.Index] {
      if seen[vertex.Index] == false {
        seen[vertex.Index] = true
        newFront := []Vertex{vertex}
        vertexStack = append(newFront, vertexStack[:]...)
      }
    }
  }
}

/*
  Runs the topological sort algorithm from the given start vertex and returns
  true if algorithm is successful.  Algorithm will fail if the graph is not
  a DAG
*/
func (g *Graph) TopoSort(start Vertex) bool {
  seen := make(map[int]bool)
  vertexQueue := []Vertex{start}
  isDag = true
  for len(vertexQueue) > 0 {
    currVertex := vertexQueue[0]
  }
  return isDag
}
