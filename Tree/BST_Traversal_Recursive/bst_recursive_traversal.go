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
	// for node,_ := range t.root {

	// }
	// fmt.Println(t)

	fmt.Println("preOder travesal of BST will be :")
	preorder(t.root)

	fmt.Println("inOder travesal of BST will be :")
	inorder(t.root)

	fmt.Println("postOder travesal of BST will be :")
	postorder(t.root)
	result := iterativepreorder(t.root)
	fmt.Println("iterative preOder travesal of BST will be :", result)
	fmt.Println("iterative full travesal of BST will be :")
}

// Rootnode->lefttreenode->righttreenode
func preorder(node *Node) {
	if node == nil {
		return
	} else {
		fmt.Println(node.data)
		preorder(node.ls)
		preorder(node.rs)
	}
}

//lefttreenode->rootnode->righttreenode

func inorder(node *Node) {
	if node == nil {
		return
	} else {
		inorder(node.ls)
		fmt.Println(node.data)
		inorder(node.rs)
	}
}

// lefttreenode->righttreenode->rootnode
func postorder(node *Node) {
	if node == nil {
		return
	} else {
		postorder(node.ls)
		postorder(node.rs)
		fmt.Println(node.data)
	}
}

func iterativepreorder(root *Node) []int {
	var res []int
	processFunc := func(v int) {
		res = append(res, v)
	}
	HelperPreOrder(root, processFunc)
	return res
}

func HelperPreOrder(node *Node, processFunc func(v int)) {
	if node != nil {
		processFunc(node.data)
		HelperPreOrder(node.ls, processFunc)
		HelperPreOrder(node.rs, processFunc)
	}
}
