package main

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func IsBalanced(root *BinaryTreeNode) bool {
	depth := 0
	return IsBalancedCore(root, &depth)
}

func IsBalancedCore(root *BinaryTreeNode, depth *int) bool {
	if root == nil {
		*depth = 0
		return true
	}
	left := 0
	right := 0
	if IsBalancedCore(root.Left, &left) && IsBalancedCore(root.Right, &right) {
		diff := left - right
		if diff <= 1 && diff >= -1 {
			if left > right {
				*depth = 1 + left
			} else {
				*depth = 1 + right
			}
			return true
		}
	}
	return false
}

func main() {
	s := &BinaryTreeNode{
		Value: 3,
	}
	fmt.Println(IsBalanced(s))
}
