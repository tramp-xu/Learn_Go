package main

type Node struct {
	Left string
	Right string
}

func (node *Node) Traverse ()  {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
}

func (node *Node) TraverseFunc(f func(*Node))  {
	if node == nill {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func main() {
	
}
