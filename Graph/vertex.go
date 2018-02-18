package graph

import (
  "fmt"
)

type Vertex struct {
  City string
  Index int
}

func (v *Vertex) String() string {
  return fmt.Sprintf("Index: %v, City: %v", v.Index, v.City)
}
