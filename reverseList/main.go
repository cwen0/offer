package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func ReverseList(root *Node) *Node {
	var reverseHead *Node
	TNode := root
	var pre *Node
	for TNode != nil {
		TNext := TNode.Next
		if TNext == nil {
			reverseHead = TNode
		}
		TNode.Next = pre
		pre = TNode
		TNode = TNext
	}
	return reverseHead
}

func main() {
	next := &Node{
		Value: 2,
	}
	head := &Node{
		Value: 1,
		Next:  next,
	}
	tHead := ReverseList(head)
	for tHead != nil {
		fmt.Println(tHead)
		tHead = tHead.Next
	}
}
