package aoc

import (
	"testing"
)

func genInts() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func genStrings() []string {
	return []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault"}
}

func genStack[T any](items []T) Stack[T] {
	stack := Stack[T]{}
	for _, v := range items {
		stack.Push(v)
	}
	return stack
}

func TestPushInts(t *testing.T) {
	items := genInts()
	stack := genStack(items)

	if len(stack.Items) != len(items) {
		t.Errorf("got %v, wanted %v", len(stack.Items), len(items))
	}

	for i, v := range items {
		if stack.Items[i] != v {
			t.Errorf("got %v, wanted %v", stack.Items[i], v)
		}
	}
}

func TestPopInts(t *testing.T) {
	items := genInts()
	stack := genStack(items)
	var expectedItem int
	var actualItem int
	var ok bool

	for i := 0; i < len(items); i++ {
		expectedItem = items[len(items)-1-i]
		actualItem, ok = stack.Pop()
		if !ok {
			t.Error("Pop should have returned true")
		}
		if actualItem != expectedItem {
			t.Errorf("got %v, wanted %v", actualItem, expectedItem)
		}
	}

	expectedItem = 0
	actualItem, ok = stack.Pop()
	if ok {
		t.Error("Popping off empty stack should return false")
	}
	if actualItem != expectedItem {
		t.Error("Value popped from empty stack should be default val for type")
	}
}

func TestPushStrings(t *testing.T) {
	items := genStrings()
	stack := genStack(items)

	if len(stack.Items) != len(items) {
		t.Errorf("got %v, wanted %v", len(stack.Items), len(items))
	}

	for i, v := range items {
		if stack.Items[i] != v {
			t.Errorf("got %v, wanted %v", stack.Items[i], v)
		}
	}
}

func TestPopStrings(t *testing.T) {
	items := genStrings()
	stack := genStack(items)
	var expectedItem string
	var actualItem string
	var ok bool

	for i := 0; i < len(items); i++ {
		expectedItem = items[len(items)-1-i]
		actualItem, ok = stack.Pop()
		if !ok {
			t.Error("Pop should have returned true")
		}
		if actualItem != expectedItem {
			t.Errorf("got %v, wanted %v", actualItem, expectedItem)
		}
	}

	expectedItem = ""
	actualItem, ok = stack.Pop()
	if ok {
		t.Error("Popping off empty stack should return false")
	}
	if actualItem != expectedItem {
		t.Error("Value popped from empty stack should be default val for type")
	}
}
