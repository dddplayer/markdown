package entity

import (
	"github.com/dddplayer/markdown/datastructure"
)

func NewTree() *tree {
	r, _ := NewRoot()
	return &tree{&datastructure.Tree{Root: r.TreeNode}}
}

type tree struct {
	*datastructure.Tree
}

func (t *tree) RootBlock() Block {
	return t.Tree.Root.Val.(*Root)
}
