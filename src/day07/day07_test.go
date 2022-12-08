package day07

import (
	"fmt"
	"testing"

	"github.com/erietz/aoc2022/src"
)

// ==============
// File Structure
// ==============
//
// ---a
//
//	|--af1
//	|--af2
//	|--ad1
//	      |---ad1f1
//	      |---ad1f2
//
// ---b
//
//	|--bf1
func generateTree() *aoc.Node[File] {
	root := aoc.NewNode(File{IsDir: true, Name: "root", Size: 0})
	a := aoc.NewNode(File{IsDir: true, Name: "a", Size: 0})
	af1 := aoc.NewNode(File{IsDir: false, Name: "af1", Size: 1})
	af2 := aoc.NewNode(File{IsDir: false, Name: "af2", Size: 1})
	ad1 := aoc.NewNode(File{IsDir: true, Name: "ad1", Size: 0})
	ad1f1 := aoc.NewNode(File{IsDir: false, Name: "ad1f1", Size: 1})
	ad1f2 := aoc.NewNode(File{IsDir: false, Name: "ad1f2", Size: 1})
	b := aoc.NewNode(File{IsDir: true, Name: "b", Size: 0})
	bf1 := aoc.NewNode(File{IsDir: false, Name: "bf1", Size: 1})

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

func TestCalcDirSize(t *testing.T) {
	root := generateTree()
	a := root.Children
	b := root.Children.Next
	ad1 := a.Children.Previous

	if a.Next != b {
		t.Error("a and b not siblings")
	}
	if a.Previous != b {
		t.Error("a and b not siblings")
	}
	if b.Previous != a {
		t.Error("a and b not siblings")
	}

	testCases := []struct {
		file *aoc.Node[File]
		size int
	}{
		{root, 5},
		{a, 4},
		{b, 1},
		{ad1, 2},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v %v", tc.file.Value.Name, tc.size), func(t *testing.T) {
			size := calcDirSize(tc.file)
			if size != tc.size {
				t.Errorf("got %v, wanted %v", size, tc.size)
			}
		})
	}

}
