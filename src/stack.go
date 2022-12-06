package aoc

import "fmt"

type Stack[T any] struct {
	Items []T
}

func (s Stack[T]) String() string {
	tmp := "{"
	for _, v := range s.Items {
		tmp += fmt.Sprintf(" %v", v)
	}
	tmp += " }"
	return tmp
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Items) == 0 {
		var tmp T
		return tmp, false
	}
	idx := len(s.Items) - 1
	top := s.Items[idx]
	s.Items = s.Items[:idx]
	return top, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.Items) == 0 {
		var tmp T
		return tmp, false
	}
	top := s.Items[len(s.Items)-1]
	return top, true
}
