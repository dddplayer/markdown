package entity

import (
	"github.com/dddplayer/markdown/datastructure"
)

func NewTree() *tree {
	return &tree{&datastructure.Tree{Root: NewNode().TreeNode}}
}

type tree struct {
	*datastructure.Tree
}

func (t *tree) RootBlock() Block {
	return t.Tree.Root.Val.(*blockNode).MdBlock
}

type blockNode struct {
	MdBlock Block
	*datastructure.TreeNode
}

func NewNode() *blockNode {
	dsn := &datastructure.TreeNode{
		Val:        nil,
		FirstChild: nil,
		LastChild:  nil,
		Parent:     nil,
		Next:       nil,
	}
	bb := &BaseBlock{state: Closed}
	n := &blockNode{
		MdBlock:  bb,
		TreeNode: dsn,
	}
	dsn.Val = n
	bb.blockNode = n

	return n
}
