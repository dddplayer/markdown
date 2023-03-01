package entity

import (
	"github.com/dddplayer/markdown/datastructure"
)

func NewTree() *blockTree {
	r, _ := NewRoot()
	return &blockTree{&datastructure.Tree{Root: r.TreeNode}}
}

type blockTree struct {
	*datastructure.Tree
}

func (t *blockTree) RootBlock() Block {
	return t.Tree.Root.Val.(*Root)
}
