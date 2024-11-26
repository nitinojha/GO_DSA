package main

import (
	"fmt"
)

type Node struct {
	data int
	ls   *Node
	rs   *Node
}

type BST struct {
	root *Node
}

// BST tree
func (tree *BST) insert(val int) {
	if tree.root == nil {
		tree.root = &Node{data: val}
	} else {
		tree.root.insert(val)
	}
}

// node
func (node *Node) insert(val int) {
	if val <= node.data {
		if node.ls == nil {
			node.ls = &Node{data: val}
		} else {
			node.ls.insert(val)
		}

	} else {
		if node.rs == nil {
			node.rs = &Node{data: val}
		} else {
			node.rs.insert(val)
		}
	}
}

func (node *Node) delete(val int) {
	if val <= node.data {
		if node.ls == nil {
			node.ls = &Node{data: val}
		} else {
			node.ls.insert(val)
		}

	} else {
		if node.rs == nil {
			node.rs = &Node{data: val}
		} else {
			node.rs.insert(val)
		}
	}
}

func main() {
	t := BST{}

	t.insert(6)
	t.insert(4)

	t.insert(5)
	t.insert(7)

	t.insert(3)
	t.insert(1)

	t.insert(8)
	t.insert(10)

	t.insert(9)
	t.insert(2)
	t.insert(0)
	//t.delete(0)

	// fmt.Println("Item Insertion order diagram is: ")

	fmt.Println("iterative full travesal of BST will be :")
	fullIterative(t.root)

}

func fullIterative(node *Node) {
	// Create a stack and start traversal at root
	stack := []*Node{}
	current := node
	// as long as current is not nil or stack is not empty, we have nodes to visit
	for current != nil || len(stack) > 0 {
		if current != nil {
			// Visit current
			fmt.Printf("%v ", current.data)
			// Push it to stack
			stack = append(stack, current)
			// Follow left link
			current = current.ls
		} else {
			// Pop stack
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// Follow right link
			current = current.rs
		}
	}
}
