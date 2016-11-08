package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func TreeDepth(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	left := TreeDepth(root.Left)
	right := TreeDepth(root.Left)
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	s := &BinaryTreeNode{
		Value: 2,
	}
	sum := TreeDepth(s)
	fmt.Println(sum)
}
