package valueobject

import (
	"github.com/dddplayer/markdown/datastructure/entity"
)

func NewTree(b Block) *BlockTree {
	return &BlockTree{&entity.Tree{Root: b.Node()}}
}

type BlockTree struct {
	*entity.Tree
}

func (t *BlockTree) RootBlock() Block {
	return t.Tree.Root.Val.(Block)
}
