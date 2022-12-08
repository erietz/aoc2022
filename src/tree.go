// adapted from https://github.com/blazingorb/ntreego/blob/master/ntree.go
package aoc

import (
	// "fmt"
	// "encoding/json"
)

type Node[T any] struct {
	Value    T
	Parent   *Node[T]
	Children *Node[T]
	Next     *Node[T]
	Previous *Node[T]
}

// func (n *Node[T]) String() string {
// 	bytes, _ := json.Marshal(n)
// 	return string(bytes)
// }

// func (n *Node[T]) String() string {
// 	s := fmt.Sprintf(
// 		"{Value: %v, Parent: %p, Children: %p, Next: %p, Previous: %p}",
// 		n.Value,
// 		n.Parent,
// 		n.Children,
// 		n.Next,
// 		n.Previous,
// 	)
// 	return s
// }

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func (n *Node[T]) AppendChild(child *Node[T]) *Node[T] {
	child.Parent = n
	if n.Children != nil {
		sibling := n.Children
		sibling.Previous = child
		for sibling.Next != nil {
			sibling = sibling.Next
		}
		child.Previous = sibling
		sibling.Next = child
	} else {
		n.Children = child
	}
	return child
}
