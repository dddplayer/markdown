package datastructure

type Tree struct {
	Root *TreeNode
}

type TreeNode struct {
	Val        any
	FirstChild *TreeNode
	LastChild  *TreeNode
	Parent     *TreeNode
	Next       *TreeNode
}

func (n *TreeNode) AppendChild(child *TreeNode) {
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
