package collections

import (
  "fmt"
  "strconv"
)

type Queue struct {
  size int
  elements []int
}

func (q *Queue) Enqueue(val int) {
  q.elements = append(q.elements, val)
  q.size++
}

func (q *Queue) Dequeue() (int, error) {
  if q.IsEmpty() {
    return 0, fmt.Errorf("Queue is empty")
  }
  first := q.elements[0]
  q.elements = append(q.elements[:1], q.elements[2:]...)
  q.size--
  return first, nil
}

func (q *Queue) IsEmpty() bool {
  return q.size == 0
}

func (q *Queue) Size() int {
  return q.size
}

func (q *Queue) String() string {
  res := "[" + strconv.Itoa(q.elements[0])
  for i := 1; i < q.size; i++ {
    res += ", " + strconv.Itoa(q.elements[i])
  }
  res += "]"
  return res
}

func (q *Queue) Peek() int {
  return q.elements[0]
}
