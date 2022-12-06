package aoc

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s Stack[T]) String() string {
	tmp := "{"
	for _, v := range s.items {
		tmp += fmt.Sprintf(" %v", v)
	}
	tmp += " }"
	return tmp
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var tmp T
		return tmp, false
	}
	idx := len(s.items) - 1
	top := s.items[idx]
	s.items = s.items[:idx]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var tmp T
		return tmp, false
	}
	top := s.items[len(s.items)-1]
	return top, true
}
