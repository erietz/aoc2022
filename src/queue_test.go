package aoc

import (
	"testing"
)

func genQueue[T any](items []T) Queue[T] {
	queue := Queue[T]{}
	for _, v := range items {
		queue.Enqueue(v)
	}
	return queue
}

func TestEnqueue(t *testing.T) {
	items := genInts()
	queue := genQueue(items)

	if len(queue.items) != len(items) {
		t.Errorf("got %v, wanted %v", len(queue.items), len(items))
	}

	for i, v := range items {
		if queue.items[i] != v {
			t.Errorf("got %v, wanted %v", queue.items[i], v)
		}
	}
}

func TestDeque(t *testing.T) {
	items := genInts()
	queue := genQueue(items)
	var expectedItem int
	var actualItem int
	var ok bool

	for i := 0; i < len(items); i++ {
		expectedItem = items[i]
		actualItem, ok = queue.Deque()
		if !ok {
			t.Error("Deque should have returned true")
		}
		if actualItem != expectedItem {
			t.Errorf("got %v, wanted %v", actualItem, expectedItem)
		}
	}

	expectedItem = 0
	actualItem, ok = queue.Deque()
	if ok {
		t.Error("Deque empty queue should return false")
	}
	if actualItem != expectedItem {
		t.Error("Value dequed from empty queue should be default val for type")
	}
}

func TestPeekEmptyQueue(t *testing.T) {
	queue := Queue[string]{}

	item, ok := queue.Peek()

	if ok {
		t.Errorf("got %v, wanted %v", ok, false)
	}

	if item != "" {
		t.Errorf("got %v, wanted %v", item, "")
	}
}
