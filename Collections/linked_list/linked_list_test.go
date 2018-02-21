/*
  Test suite on the linked list Implementation
*/
package collections

import (
  "testing"
)

/*
  Tests adding to an empty linked list
*/
func TestAddToEmptyList(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  if list.Size != 1 {
    t.Errorf("Expected size %v, Got %v", 1, list.Size)
  }
}

/*
  Tests adding to an already partially filled in linked list
*/
func TestAddToPartially(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  list.Add(2)
  list.Add(3)
  if list.Size != 3 {
    t.Errorf("Expected size %v, Got %v", 3, list.Size)
  }
  val, err := list.Get(2)
  if val != 3 && err == nil {
    t.Errorf("Expected %v, Got %v", 3, val)
  }
}

// Tests adding to index 0 when list is empty
func TestAddAtIndexEmpty(t *testing.T) {
  list := new(LinkedList)
  list.AddAtIndex(1, 0)
  if list.Size != 1 {
    t.Errorf("Expected size %v, Got %v", 1, list.Size)
  }
  if val, err := list.Get(0); val != 1 || err != nil {
    t.Errorf("Expected front value to be %v, Got %v", 1, val)
  }
}

// Tests adding at an index
func TestAddAtIndex(t *testing.T) {
  expected := []int{1, 2, 3, 4}
  list := new(LinkedList)
  list.Add(1)
  list.Add(2)
  list.Add(4)
  list.AddAtIndex(3, 2)
  for i := 0; i < list.Size; i++ {
    val, err := list.Get(i)
    if err != nil || val != expected[i] {
      t.Errorf("Numbers don't match.  Expected %v, Got %v", list, expected)
    }
  }
}

// Tests adding at an index larger than size of list
func TestAddOutOfBounds(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  list.Add(2)
  list.Add(3)
  err := list.AddAtIndex(1, 10)
  if err == nil {
    t.Errorf("Error not found")
  }
}

/*
  Tests get on an empty list
*/
func TestGetOnEmptyList(t *testing.T) {
  list := new(LinkedList)
  _, err := list.Get(1)
  if err == nil {
    t.Error("Expected error but none was given")
  }
}

/*
  Tests get when the supplied parameter is out of the range
*/
func TestOutOfBoundsGet(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  _, err := list.Get(2)
  if err == nil {
    t.Error("Expected error but none was given")
  }
}

/*
  Tests get on a filled list
*/
func TestGet(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  value, err := list.Get(0)
  if value != 1 && err == nil {
    t.Errorf("Expected %v, Got %v", 1, value)
  }
}

/*
  Tests contains on an empty list
*/
func TestContainsOnEmpty(t *testing.T) {
  list := new(LinkedList)
  res := list.Contains(1)
  if res {
    t.Error("Expected false got", res)
  }
}

/*
  Tests contains when element not present in list
*/
func TestContainsWhenNotPresent(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  res := list.Contains(3)
  if res {
    t.Error("Expected false got", res)
  }
}

/*
  Tests contains when element is present in list
*/
func TestContainsWhenPresent(t *testing.T) {
  list := new(LinkedList)
  list.Add(1)
  res := list.Contains(1)
  if !res {
    t.Error("Expected true got", res)
  }
}
