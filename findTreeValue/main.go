package main

import "fmt"

type Node struct {
	Value int
	Flag  bool
	Left  *Node
	Right *Node
}

type NodeStack []Node

func (n *NodeStack) Pop() {
	*n = (*n)[:len(*n)-1]
}
func (n *NodeStack) Top() Node {
	return (*n)[len(*n)-1]
}
func (n *NodeStack) Size() int {
	return len(*n)
}
func (n *NodeStack) Push(node Node) {
	node.Flag = true
	*n = append(*n, node)
}

func preOrder(root *Node) {
	var stack NodeStack
	// 	stack.Push(*root)
	//for stack.Size() > 0 {
	//node := stack.Top()
	//stack.Pop()
	//fmt.Printf("%d ", node.Value)
	//if node.Right != nil {
	//stack.Push(*(node.Right))
	//}
	//if node.Left != nil {
	//stack.Push(*(node.Left))
	//}
	//}
	node := root
	for node != nil || stack.Size() > 0 {
		for node != nil {
			fmt.Printf("%d ", node.Value)
			//if node.Value == 7 {
			//fmt.Printf("%d ", node.Value)
			//for stack.Size() > 0 {
			//t := stack.Top()
			//stack.Pop()
			//fmt.Printf("%d ", t.Value)
			//}
			//}
			stack.Push(*node)
			node = (*node).Left
		}
		if stack.Size() > 0 {
			t := stack.Top()
			stack.Pop()
			node = t.Right
		}
	}
}

func main() {
	left := Node{
		Value: 4,
		Left: &Node{
			Value: 5,
		},
		Right: &Node{
			Value: 6,
		},
	}
	right := Node{
		Value: 9,
		Left: &Node{
			Value: 7,
		},
		Right: &Node{
			Value: 10,
		},
	}
	root := Node{
		Value: 3,
		Left:  &left,
		Right: &right,
	}
	preOrder(&root)
	fmt.Println()
}
