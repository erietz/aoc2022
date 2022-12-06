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

func TestPeek(t *testing.T) {
	type foo struct {
		t1 int
		t2 string
		t3 bool
	}
	f1 := foo{t1: 1, t2: "f1", t3: true}
	f2 := foo{t1: 2, t2: "f2", t3: false}
	f3 := foo{t1: 3, t2: "f3", t3: true}
	stack := Stack[foo]{}
	stack.Push(f1)
	stack.Push(f2)
	stack.Push(f3)

	last, ok := stack.Peek()
	if ok != true {
		t.Errorf("got %v, wanted %v", ok, true)
	}
	if len(stack.Items) != 3 {
		t.Errorf("got %v, wanted %v", len(stack.Items), 3)
	}
	if last != f3 {
		t.Errorf("got %v, wanted %v", last, f3)
	}

	last.t2 = "last and f3 are separate copies of type foo"
	if last == f3 {
		t.Errorf("got %v, wanted %v", last, f3)
	}
}
