package day07

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/erietz/aoc2022/src"
)

type File struct {
	Name  string
	Size  int
	IsDir bool
}

func Solve(input string) {
	root := parseTree(input)
	calcDirSizes(root.Children)
	fmt.Println(sumDirsLessThan100000(root.Children))
}

func sumDirsLessThan100000(root *aoc.Node[File]) int {
	return _sumDirsLessThan100000(root, 0)
}

func _sumDirsLessThan100000(root *aoc.Node[File], sum int) int {
	if root == nil {
		return sum
	}

	if root.Value.IsDir && root.Value.Size < 100000 {
		sum += root.Value.Size
	}

	return _sumDirsLessThan100000(root.Children, sum) + _sumDirsLessThan100000(root.Next, sum)
}

func calcDirSizes(root *aoc.Node[File]) {
	if root == nil {
		return
	}

	if root.Value.IsDir {
		root.Value.Size = calcDirSize(root)
	}

	calcDirSizes(root.Children)
	calcDirSizes(root.Next)
}

func calcDirSize(file *aoc.Node[File]) int {
	return _calcDirSize(file.Children, 0)
}

func _calcDirSize(file *aoc.Node[File], sum int) int {
	if file == nil {
		return sum
	}
	sum += file.Value.Size
	return _calcDirSize(file.Children, sum)
}


func traverseTree(root *aoc.Node[File]) {
	if root == nil {
		return
	}

	fmt.Println(root.Value.Name)
	traverseTree(root.Children)
	traverseTree(root.Next)
}

func parseTree(input string) *aoc.Node[File] {
	curr := aoc.NewNode(File{Name: "root", IsDir: true, Size: 0})
	root := curr

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		if strings.HasPrefix(line, "$ ") {
			cmdParts := strings.Split(line[2:], " ")

			switch cmdParts[0] {
			case "cd":
				name := cmdParts[1]
				if name == ".." {
					curr = curr.Parent
				} else {
					child := aoc.NewNode(File{Name: name, IsDir: true, Size: 0})
					curr.AppendChild(child)
					curr = child
				}
			case "ls":
				continue
			}
		} else {
			cmdParts := strings.Split(line, " ")
			switch line {
			case "dir":
				name := cmdParts[1]
				child := File{Name: name, IsDir: true, Size: 0}
				curr.AppendChild(aoc.NewNode(child))
			default:
				if size, err := strconv.Atoi(cmdParts[0]); err == nil {
					name := cmdParts[1]
					child := File{Name: name, IsDir: false, Size: size}
					curr.AppendChild(aoc.NewNode(child))
				}
			}
		}
	}
	return root
}
