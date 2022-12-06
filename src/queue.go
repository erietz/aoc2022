package aoc

import (
	"fmt"
)

type Queue[T any] struct {
	items []T
}

func (q Queue[T]) String() string {
	tmp := "{"
	for _, v := range q.items {
		tmp += fmt.Sprintf(" %v", v)
	}
	tmp += " }"
	return tmp
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Deque() (T, bool) {
	if len(q.items) == 0 {
		var tmp T
		return tmp, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var tmp T
		return tmp, false
	}
	return q.items[0], true
}
