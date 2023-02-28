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

type WalkStatus int

const (
	WalkStop WalkStatus = iota + 1
	WalkContinue
)

type WalkState int

const (
	WalkIn WalkState = 1 << iota
	WalkOut
)

type Walker func(v any, ws WalkState) WalkStatus

func (t *Tree) Walk(walker Walker) {
	walkNode(t.Root, walker)
}

func walkNode(n *TreeNode, walker Walker) WalkStatus {
	status := walker(n.Val, WalkIn)
	if status != WalkStop {
		for c := n.FirstChild; c != nil; c = c.Next {
			if s := walkNode(c, walker); s == WalkStop {
				return WalkStop
			}
		}
	}
	return walker(n.Val, WalkOut)
}
