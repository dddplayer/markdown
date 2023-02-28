package datastructure

type Tree struct {
	Root *Node
}

type Node struct {
	Val        any
	FirstChild *Node
	LastChild  *Node
	Parent     *Node
	Next       *Node
}

func (n *Node) AppendChild(child *Node) {
	if n.FirstChild == nil {
		n.FirstChild = child
		child.Next = nil
	} else {
		last := n.LastChild
		last.Next = child
	}
	child.Parent = n
	n.LastChild = child
}
