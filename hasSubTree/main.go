package main

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func HasSubTree(root1, root2 *Node) bool {
	var result bool
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}

	if root1.Value == root2.Value {
		result = DoesTree1HaveTree2(root1, root2)
	}
	if !result {
		HasSubTree(root1.Left, root2)
	}
	if !result {
		HasSubTree(root1.Right, root2)
	}
	return result
}

func DoesTree1HaveTree2(root1, root2 *Node) bool {
	if root2 == nil {
		return true
	}

	if root1 == nil {
		return false
	}

	if root1.Value != root2.Value {
		return false
	}
	return DoesTree1HaveTree2(root1.Left, root2.Left) && DoesTree1HaveTree2(root1.Right, root2.Right)
}

func main() {

}
