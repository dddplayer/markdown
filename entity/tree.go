package entity

import "github.com/dddplayer/markdown/datastructure"

func NewTree() *tree {
	return &tree{&datastructure.Tree{Root: NewNode().Node}}
}

type tree struct {
	*datastructure.Tree
}

func (t *tree) RootBlock() Block {
	return t.Tree.Root.Val.(*node).B
}

type node struct {
	B Block
	*datastructure.Node
}

func NewNode() *node {
	dsn := &datastructure.Node{
		Val:        nil,
		FirstChild: nil,
		LastChild:  nil,
		Parent:     nil,
		Next:       nil,
	}
	n := &node{
		B:    nil,
		Node: dsn,
	}
	dsn.Val = n

	return n
}

func (n *node) AppendBlock(b Block) {
	n.Node.AppendChild(b.Node().Node)
}
