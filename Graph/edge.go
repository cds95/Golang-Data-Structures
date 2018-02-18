package graph

import (
  "fmt"
)

type Edge struct {
  Source Vertex
  Destination Vertex
  Weight int
}

func (e *Edge) ToString() string {
  return fmt.Sprintf("From: %v To: %v Weight: %v", e.Source, e.Destination, e.Weight)
}
