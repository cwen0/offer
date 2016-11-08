package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type NodeStack []Node

func (n *NodeStack) Pop() {
	*n = (*n)[:len(*n)-1]
}

func (n *NodeStack) Top() Node {
	//if n.Size() == 1 {
	//return
	//}
	return (*n)[len(*n)-1]
}

func (n *NodeStack) Size() int {
	return len(*n)
}

func (n *NodeStack) Push(node Node) {
	*n = append(*n, node)
}

func preOrder(root *Node) {
	var stack NodeStack
	stack.Push(*root)
	for stack.Size() > 0 {
		node := stack.Top()
		stack.Pop()
		fmt.Printf("%d ", node.Value)
		if node.Right != nil {
			stack.Push(*(node.Right))
		}
		if node.Left != nil {
			stack.Push(*(node.Left))
		}
	}
}

func inOrder(root *Node) {
	var stack NodeStack
	p := root
	for p != nil || stack.Size() > 0 {
		if p != nil {
			stack.Push(*p)
			p = p.Left
		} else {
			t := stack.Top()
			p = &t
			stack.Pop()
			fmt.Printf("%d ", p.Value)
			p = p.Right
		}
	}
}

func postOrder(root *Node) {
	var stack1 NodeStack
	var stack2 NodeStack
	stack1.Push(*root)
	for stack1.Size() > 0 {
		node := stack1.Top()
		stack1.Pop()
		stack2.Push(node)
		if node.Left != nil {
			stack1.Push(*(node.Left))
		}
		if node.Right != nil {
			stack1.Push(*(node.Right))
		}
	}

	for stack2.Size() > 0 {
		fmt.Printf("%d ", stack2.Top().Value)
		stack2.Pop()
	}
}

func main() {
	left := Node{Value: 4}
	right := Node{Value: 9}
	root := Node{
		Value: 3,
		Left:  &left,
		Right: &right,
	}
	preOrder(&root)
	fmt.Println()
	inOrder(&root)
	fmt.Println()
	postOrder(&root)
}
