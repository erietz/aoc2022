package aoc

import (
	"testing"
)

type file struct {
	name string
	size int
	isDir bool
}

// ==============
// File Structure
// ==============
//
// ---a
//    |--af1
//    |--af2
//    |--ad1
//          |---ad1f1
//          |---ad1f2
// ---b
//    |--bf1
func generateTree() *Node[file] {
	root := NewNode(file{isDir: true, name: "root"})
	a := NewNode(file{isDir: true, name: "a"})
	af1 := NewNode(file{isDir: false, name: "af1"})
	af2 := NewNode(file{isDir: false, name: "af2"})
	ad1 := NewNode(file{isDir: true, name: "ad1"})
	ad1f1 := NewNode(file{isDir: false, name: "ad1f1"})
	ad1f2 := NewNode(file{isDir: false, name: "ad1f2"})
	b := NewNode(file{isDir: true, name: "b"})
	bf1 := NewNode(file{isDir: false, name: "bf1"})

	root.AppendChild(a)
	a.AppendChild(af1)
	a.AppendChild(af2)
	a.AppendChild(ad1)
	ad1.AppendChild(ad1f1)
	ad1.AppendChild(ad1f2)
	root.AppendChild(b)
	b.AppendChild(bf1)
	return root
}

func TestAppendChild(t *testing.T) {
	root := generateTree()
	a := root.Children
	af1 := a.Children
	af2 := af1.Next
	ad1 := af2.Next
	ad1f1 := ad1.Children
	ad1f2 := ad1f1.Next
	b := a.Next
	bf1 := b.Children

	if root.Parent != nil {
		t.Error("root.Parent should be nil")
	}
	if root.Value.name != "root" {
		t.Errorf("got %v, wanted %v", root.Value.name, "root")
	}
	if a.Value.name != "a" {
		t.Errorf("got %v, wanted %v", a.Value.name, "a")
	}
	if af1.Value.name != "af1" {
		t.Errorf("got %v, wanted %v", af1.Value.name, "af1")
	}
	if af2.Value.name != "af2" {
		t.Errorf("got %v, wanted %v", af2.Value.name, "af2")
	}
	if ad1.Value.name != "ad1" {
		t.Errorf("got %v, wanted %v", ad1.Value.name, "ad1")
	}
	if ad1f1.Value.name != "ad1f1" {
		t.Errorf("got %v, wanted %v", ad1f1.Value.name, "ad1f1")
	}
	if ad1f2.Value.name != "ad1f2" {
		t.Errorf("got %v, wanted %v", ad1f2.Value.name, "ad1f2")
	}
	if b.Value.name != "b" {
		t.Errorf("got %v, wanted %v", b.Value.name, "b")
	}
	if bf1.Value.name != "bf1" {
		t.Errorf("got %v, wanted %v", bf1.Value.name, "bf1")
	}
}
