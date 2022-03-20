package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TODO: after the implementation use generics (go 1.18).
// TODO: reverse the binary tree without recursion.
// TODO: dump the tree as a graphviz graph.

// Node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func generateTree(depth, value int, root *Node) *Node {
	if depth == 0 {
		return nil
	}

	rand.Seed(time.Now().UnixNano())

	v := rand.Intn(100)

	if root == nil {
		root = &Node{
			Value: value,
		}
		depth -= 1
		value -= v

		root.Left = generateTree(depth, (value+v)+(v+2), root.Left)
		root.Right = generateTree(depth, (value+v)+(v+5), root.Right)
	}

	return root
}

// PrintTree
func (n Node) PrintTree(root *Node, d int) {
	if root == nil {
		return
	}
	d += d
	space := ""
	for i := 0; i < d; i++ {
		space += " "
	}
	n.PrintTree(root.Left, d)
	fmt.Printf("%v%v\n", space, root.Value)
	n.PrintTree(root.Right, d)
}

// ReverseTree
func ReverseTree(root *Node) *Node {

	if root == nil {
		return nil
	}

	tmp := &Node{Value: root.Value}
	tmp.Left = root.Right
	tmp.Right = root.Left

	root = tmp

	root.Left = ReverseTree(root.Left)
	root.Right = ReverseTree(root.Right)

	return root
}

func main() {

	const (
		depth      = 4
		startValue = 4
		space      = 2
	)

	r := generateTree(depth, startValue, nil)
	r.PrintTree(r, space)

	r = ReverseTree(r)

	fmt.Println()
	fmt.Println("------------------------------")
	fmt.Println()

	r.PrintTree(r, space)

}
