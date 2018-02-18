package main

import (
  "DataStructures/graph"

)

func main() {
  v1 := graph.Vertex{City: "Seattle", Index: 1}
  v2 := graph.Vertex{City: "New York", Index: 2}
  v3 := graph.Vertex{City: "Miami", Index: 3}
  v4 := graph.Vertex{City: "Los Angeles", Index: 4}

  e1 := graph.Edge{Source: v1, Destination: v3, Weight: 500}
  e2 := graph.Edge{Source: v1, Destination: v2, Weight: 400}
  e3 := graph.Edge{Source: v2, Destination: v4, Weight: 800}

  graph := new(graph.Graph)
  graph.AddVertex(v1)
  graph.AddVertex(v2)
  graph.AddVertex(v3)
  graph.AddVertex(v4)

  graph.AddEdge(e1)
  graph.AddEdge(e2)
  graph.AddEdge(e3)



  graph.DFS(v1)
}
