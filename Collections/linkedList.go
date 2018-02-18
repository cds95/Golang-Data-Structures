/*
  Implementation of a linked list
*/

package collections

import (
  "fmt"
  "strconv"
)

type Node struct {
  data int
  next *Node
}

type LinkedList struct {
  front *Node
  Size int
}

// Checks to see whether or not the list is empty.  Returns true if it is and
// false otherwise
func (list *LinkedList) IsEmpty() bool {
  return list.Size == 0
}

// Adds a number to the end of the list
func (list *LinkedList) Add(val int) {
  node := Node{data: val}
  if list.IsEmpty() {
    list.front = &node
  } else {
    curr := list.front
    for curr.next != nil {
      curr = curr.next
    }
    curr.next = &node
  }
  list.Size++
}

// Returns the element at the index position
func (list *LinkedList) Get(index int) (int, error){
  if index < 0 || index > list.Size {
    return -1, fmt.Errorf("Index out of bound exception")
  }
  curr := list.front
  for i := 0; i < index; i++ {
    curr = curr.next
  }
  return curr.data, nil
}

// Adds an number at the specified index of the linked list.
// Returns true if successful and false otherwise.  It will fail
// if the given index is negative or more than the size
func (list *LinkedList) AddAtIndex(val int, index int) bool {
  if index < 0 || index > list.Size {
    return false
  }
  node := Node{data: val}
  curr := list.front
  for i := 0; i < index - 1; i++ {
    curr = curr.next
  }
  temp := curr.next
  curr.next = &node
  curr.next.next = temp
  list.Size++
  return true
}

// Prints out the list
func (list *LinkedList) Print() {
  fmt.Print("[")
  if !list.IsEmpty() {
    fmt.Print(list.front.data)
    curr := list.front
    for curr.next != nil {
      fmt.Printf(", %v", curr.next.data)
      curr = curr.next
    }

  }
  fmt.Println("]")
}

// Linked list to string method
func (list *LinkedList) String() string {
  result := "["
  if !list.IsEmpty() {
    result += strconv.Itoa(list.front.data)
    curr := list.front
    for curr.next != nil {
      result += ", " + strconv.Itoa(curr.next.data)
      curr = curr.next
    }

  }
  result += "]"
  return result
}

// Checks to see whether or not the list contains a given number
// Returns true if it contains it and false otherwise
func (list *LinkedList) Contains(val int) bool {
  if list.IsEmpty() {
    return false
  }
  curr := list.front
  for curr != nil {
    if curr.data == val {
      return true
    } else {
      curr = curr.next
    }
  }
  return false
}

// Removes a number from the linked list
func (list *LinkedList) Delete(val int) bool {
  for list.front.data == val {
    list.front = list.front.next
  }
  curr := list.front
  for curr.next != nil {
    if curr.next.data == val {
      curr.next = curr.next.next
      return true
    } else {
      curr = curr.next
    }
  }
  return false
}
