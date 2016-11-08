package main

type ListNode struct {
	Key  int
	Next *ListNode
}

func FindFirstCommentNode(head1, head2 *ListNode) *ListNode {
	lengthOne := GetListNodeLength(head1)
	lengthTwo := GetListNodeLength(head2)
	lengthDif := lengthOne - lengthTwo
	listLong := head1
	listShort := head2
	if lengthOne < lengthTwo {
		lengthDif = lengthTwo - lengthOne
		listLong = head2
		listShort = head1
	}

	for i := 0; i < lengthDif; i++ {
		listLong = listLong.Next
	}

	for listLong != nil && listShort != nil && listLong != listShort {
		listLong = listLong.Next
		listShort = listShort.Next
	}
	return listShort
}

func GetListNodeLength(head *ListNode) int {
	count := 1
	for {
		if head.Next == nil {
			break
		}
		count++
		head = head.Next
	}
	return count
}

func main() {
}
