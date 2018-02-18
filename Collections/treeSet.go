package collections

import (
  "strconv"
)

type TreeNode struct {
  val int
  left *TreeNode
  right *TreeNode
}

type TreeSet struct {
  root *TreeNode
  size int
}

/*
  Returns the current size of the set
*/
func (t *TreeSet) Size() int {
  return t.size
}

/*
  Checks to see whether the tree set is empty.  Returns true if its is empty
  and false otherwise
*/
func (t *TreeSet) IsEmpty() bool {
  return t.root == nil
}

/*
  Adds a given value into the tree set.  Returns true if adding was
  successful and false otherwise.  This will return false if the set already
  contains the paremeter value
*/
func (t *TreeSet) Add(val int) bool {
  if t.Contains(val) {
    return false
  }
  t.root = addHelper(val, t.root)
  t.size++
  return true
}

func addHelper(data int, curr *TreeNode) *TreeNode {
  if curr == nil {
    newNode := TreeNode{val: data}
    return &newNode
  }
  if data < curr.val {
    curr.left = addHelper(data, curr.left)
  } else {
    curr.right = addHelper(data, curr.right)
  }
  return curr
}

/*
  Checks to see whether or not set contains a given value. Returns true if
  value is present and false otherwise
*/
func (t *TreeSet) Contains(val int) bool {
  return containsHelper(val, t.root)
}

func containsHelper(data int, curr *TreeNode) bool {
  if curr == nil {
    return false
  }
  if curr.val == data {
    return true
  } else if data < curr.val {
    return containsHelper(data, curr.left)
  } else {
    return containsHelper(data, curr.right)
  }
}

/*
  Returns the string representation of a tree set.  Elements are printed out
  according to a in order tranversal
*/
func (t* TreeSet) String() string {
  return stringHelper(t.root)
}

func stringHelper(curr *TreeNode) string {
  if (curr != nil) {
    return stringHelper(curr.left) + " " + strconv.Itoa(curr.val) + stringHelper(curr.right)
  }
  return ""
}
