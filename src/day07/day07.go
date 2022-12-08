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
	// Part 1: 1749646
	fmt.Println("Part 1:", sumDirsLessThanSize(root, 100000))
	// Part 2: 1498966
	fmt.Println("Part 2:", findDirToDelete(root))
}

func sumDirsLessThanSize(root *aoc.Node[File], size int) int {
	sum := 0
	return sumDirsLessThanSizeHelper(root, size, &sum)
}

func findDirToDelete(root *aoc.Node[File]) int {
	totalDiskSize := 70000000
	updateSize := 30000000
	rootSize := calcDirSize(root)
	freeSpace := totalDiskSize - rootSize
	spaceNeededForUpdate := updateSize - freeSpace

	tmp := totalDiskSize
	smallestDir := &tmp
	return findDirToDeleteHelper(root, spaceNeededForUpdate, smallestDir)
}

func findDirToDeleteHelper(root *aoc.Node[File], spaceNeededForUpdate int, smallestDir *int) int {
	if root == nil {
		return 0
	}

	if root.Value.IsDir {
		dirSize := calcDirSize(root)
		if dirSize >= spaceNeededForUpdate && dirSize <= *smallestDir {
			*smallestDir = dirSize
		}
	}

	for curr := root.Children; curr != nil; curr = curr.Next {
		findDirToDeleteHelper(curr, spaceNeededForUpdate, smallestDir)
	}

	return *smallestDir
}

func sumDirsLessThanSizeHelper(root *aoc.Node[File], size int, sum *int) int {
	if root == nil {
		return 0
	}

	if root.Value.IsDir {
		dirSize := calcDirSize(root)
		if dirSize < size {
			*sum += dirSize
		}
	}

	for curr := root.Children; curr != nil; curr = curr.Next {
		sumDirsLessThanSizeHelper(curr, size, sum)
	}

	return *sum
}

func calcDirSize(file *aoc.Node[File]) int {
	if !file.Value.IsDir {
		panic("File is not a directory")
	}

	size := 0
	return calcDirSizeHelper(file, &size)
}

func calcDirSizeHelper(file *aoc.Node[File], sum *int) int {
	if file == nil {
		return 0
	}

	if !file.Value.IsDir {
		*sum += file.Value.Size
	}

	for curr := file.Children; curr != nil; curr = curr.Next {
		calcDirSizeHelper(curr, sum)
	}
	return *sum
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
