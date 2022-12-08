package aoc

type Node[T any] struct {
	Value    T
	Parent   *Node[T]
	Children *Node[T]
	Next     *Node[T]
	Previous *Node[T]
}

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
