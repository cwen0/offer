package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func Merge(head1, head2 *Node) *Node {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var mergeHead *Node

	if head1.Value < head2.Value {
		mergeHead = head1
		mergeHead.Next = Merge(head1.Next, head2)
	} else {
		mergeHead = head2
		mergeHead.Next = Merge(head1, head2.Next)
	}
	return mergeHead
}

func main() {
	head1 := &Node{
		Value: 1,
		Next: &Node{
			Value: 3,
		},
	}

	head2 := &Node{
		Value: 2,
		Next: &Node{
			Value: 4,
		},
	}

	tHead := Merge(head1, head2)
	for tHead != nil {
		fmt.Println(tHead.Value)
		tHead = tHead.Next
	}
}
