package main

import "fmt"

//BST used to construct more abstract data structures such as sets, multisets, and associative arrays (maps, multimaps, etc.).
type BST struct {
	root *Node
}

type Node struct {
	Data  int
	left  *Node
	right *Node
}

func (node *Node) insert(val int) {
	if val <= node.Data {
		if node.left == nil {
			node.left = &Node{Data: val}
		} else {
			node.left.insert(val)
		}
	} else {
		if node.right == nil {
			node.right = &Node{Data: val}
		} else {
			node.right.insert(val)
		}
	}

}

func (t *BST) insert(val int) {
	node := &Node{Data: val}
	if t.root == nil {
		t.root = node
	} else {
		//node := &Node{Data : val}
		t.root.insert(val)
	}

}

func main() {
	t := BST{}
	t.insert(10)
	t.insert(6)
	t.insert(8)
	t.insert(7)

	t.insert(11)
	t.insert(12)
	t.insert(13)
	t.insert(15)
	t.insert(14)

	fmt.Println("Item Insertion order diagram is: ")
	// for node,_ := range t.root {

	// }
	fmt.Println(t.root, t.root.left, t.root.left.right, t.root.right, t.root.right.left, t.root.right.right, t.root.right.right.right)

}
