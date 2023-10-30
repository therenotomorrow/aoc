// Package day07: https://adventofcode.com/2022/day/7
package day07

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tree struct {
	Root *Node `json:"root"`
}

type Node struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	IsDir    bool   `json:"is_dir"`
	parent   *Node
	Children []*Node `json:"children,omitempty"`
}

func (n *Node) findDir(dir string) *Node {
	for _, child := range n.Children {
		if child.Name == dir && child.IsDir {
			return child
		}
	}

	return nil
}

func (n *Node) calcDirSize() {
	total := 0

	for _, child := range n.Children {
		if child.IsDir {
			child.calcDirSize()
		}
		total += child.Size
	}

	n.Size = total
}

func (n *Node) Sizes() []int {
	s := make([]int, 0)

	if n.IsDir {
		s = append(s, n.Size)
	}

	for _, child := range n.Children {
		if child.IsDir {
			s = append(s, child.Sizes()...)
		}
	}

	return s
}

func NewTree() Tree {
	return Tree{&Node{Name: "/", IsDir: true}}
}

func CollectTree(actions []string) Tree {
	tree := NewTree()
	var curr *Node
	var wait bool

	for _, action := range actions {
		switch {

		case strings.HasPrefix(action, "$ cd "):
			switch dir := strings.TrimPrefix(action, "$ cd "); dir {
			case "/":
				curr = tree.Root
			case "..":
				curr = curr.parent
			default:
				curr = curr.findDir(dir)
			}
			wait = false

		case strings.HasPrefix(action, "$ ls"):
			wait = true

		case wait:
			var newNode *Node
			split := strings.Fields(action)

			switch split[0] {
			case "dir":
				newNode = &Node{Name: split[1], IsDir: true}
			default:
				size, _ := strconv.Atoi(split[0])
				newNode = &Node{Name: split[1], Size: size}
			}

			newNode.parent = curr
			curr.Children = append(curr.Children, newNode)
		}
	}

	return tree
}

// 1315285
func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()
	scan := bufio.NewScanner(f)

	var actions []string

	for scan.Scan() {
		actions = append(actions, scan.Text())
	}

	tr := CollectTree(actions)
	tr.Root.calcDirSize()
	bb, _ := json.MarshalIndent(tr, "", "  ")
	//fmt.Printf("%s", bb)
	//fmt.Println(tr.Root.Sizes())
	fs, _ := os.Create("memes.json")
	defer fs.Close()
	fs.Write(bb)

	sizes := tr.Root.Sizes()
	sort.Ints(sizes)
	// part 1
	max := 0
	for _, ss := range sizes {
		if ss <= 100000 {
			max += ss
		}
	}
	//fmt.Println(max)
	//fds := make([]int, 0)
	//
	//for i := len(sizes) - 1; i >= 0; i-- {
	//	total := sizes[i]
	//
	//	for j := len(sizes) - 1 - i; j >= 0; j-- {
	//		if total+sizes[j] <= 100000 {
	//			total += sizes[j]
	//		}
	//	}
	//	fds = append(fds, total)
	//}
	//sort.Ints(fds)
	//fmt.Println(fds)
	totalSpace := 70000000
	desireSize := 30000000
	currentSize := tr.Root.Size

	currFree := totalSpace - currentSize
	search := desireSize - currFree
	//9847279
	//min := 0
	// part 2
	var ss int
	for _, size := range sizes {
		if size >= search {
			ss = size
			break
		}
	}
	fmt.Println(ss)
}
