package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func FindKthToTail(head *Node, k int) *Node {
	if head == nil || k == 0 {
		return nil
	}

	Ahead := head
	for i := 0; i < k-1; i++ {
		if Ahead.Next != nil {
			Ahead = Ahead.Next
		} else {
			return nil
		}
	}

	Behind := head
	for Ahead.Next != nil {
		Ahead = Ahead.Next
		Behind = Behind.Next
	}
	return Behind
}

func main() {
	next := Node{
		Value: 2,
	}
	head := Node{
		Value: 1,
		Next:  &next,
	}
	result := FindKthToTail(&head, 2)
	fmt.Println(result)
}
