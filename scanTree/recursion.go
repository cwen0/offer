package main

import "fmt"

type Node struct {
	Vlaue int
	Left  *Node
	Right *Node
}

func preOrder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Vlaue)
	preOrder(root.Left)
	preOrder(root.Right)
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	fmt.Printf("%d ", root.Vlaue)
	preOrder(root.Right)
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	preOrder(root.Left)
	preOrder(root.Right)
	fmt.Printf("%d ", root.Vlaue)
}

func main() {
	left := Node{Vlaue: 4}
	right := Node{Vlaue: 9}
	root := Node{
		Vlaue: 3,
		Left:  &left,
		Right: &right,
	}
	preOrder(&root)
	fmt.Println()
	inOrder(&root)
	fmt.Println()
	postOrder(&root)
}
