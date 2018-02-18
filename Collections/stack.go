// Stack implementation using a slice

package collections

import (
  "fmt"
  "strconv"
)

type Stack struct {
  size int
  elements []int
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
func (s *Stack) Push(val int) {
  newSlice := []int{val}
  s.elements = append(newSlice, s.elements...)
  s.size++
}

// Pops a value from the top of the stack and returns it
func (s *Stack) Pop() (int, error) {
  if s.IsEmpty() {
    return 0, fmt.Errorf("Empty Stack")
  }
  first := s.elements[0]
  s.elements = append(s.elements[:0], s.elements[1:]...)
  s.size--
  return first, nil
}

// ToString method
func (s *Stack) String() string {
  res := "[" + strconv.Itoa(s.elements[0])
  for i, num := range s.elements {
    if i > 0 {
      res += ", " + strconv.Itoa(num)
    }
  }
  res += "]"
  return res
}

// Returns the top value on the stack
func (s *Stack) Peek() int {
  return s.elements[0]
}
