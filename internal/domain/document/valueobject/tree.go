package valueobject

import (
	"github.com/dddplayer/markdown/pkg/datastructure"
)

func NewTree(b Block) *BlockTree {
	return &BlockTree{&datastructure.Tree{Root: b.Node()}}
}

type BlockTree struct {
	*datastructure.Tree
}

func (t *BlockTree) RootBlock() Block {
	return t.Tree.Root.Val.(Block)
}
