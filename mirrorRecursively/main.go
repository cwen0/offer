package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func MirrorRecursively(root *Node) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		MirrorRecursively(root.Left)
	}
	if root.Right != nil {
		MirrorRecursively(root.Right)
	}
}

func main() {
	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
		},
		Right: &Node{
			Value: 3,
		},
	}
	MirrorRecursively(root)
	fmt.Println(root, root.Left, root.Right)
}
