// Stack implementation using a slice

package collections

import (
  "fmt"
)

type Stack struct {
  size int
  elements []interface{}
}

// Returns the size of the stack
func (s *Stack) Size() int {
  return s.size
}

// Checks to see whether or no the stack is empty.  Returns true if it empty
// and false otherwise
func (s *Stack) IsEmpty() bool {
  return s.size == 0
}

// Pushes a new value on to the top of the stack
func (s *Stack) Push(val interface{}) {
  var newSlice []interface{} = make([]interface{}, 1)
  newSlice[0] = val
  s.elements = append(newSlice, s.elements...)
  s.size++
}

// Pops a value from the top of the stack and returns it
func (s *Stack) Pop() (interface{}, error) {
  if s.IsEmpty() {
    return -1, fmt.Errorf("Empty Stack")
  }
  first := s.elements[0]
  s.elements = append(s.elements[:0], s.elements[1:]...)
  s.size--
  return first, nil
}

// ToString method
func (s *Stack) String() string {
  res := "[" + fmt.Sprintf("%v", s.elements[0])
  for i, val := range s.elements {
    if i > 0 {
      res += fmt.Sprintf(", %v", val)
    }
  }
  res += "]"
  return res
}

// Returns the top value on the stack
func (s *Stack) Peek() interface{} {
  return s.elements[0]
}
