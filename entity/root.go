package entity

import (
	"github.com/dddplayer/markdown/datastructure"
	"github.com/dddplayer/markdown/parser"
	"github.com/dddplayer/markdown/parser/valueobject"
)

type Root struct {
	*BaseBlock
}

func NewRoot() (*Root, error) {
	r := &Root{
		BaseBlock: &BaseBlock{
			TreeNode: datastructure.EmptyTreeNode(),
			Parser:   parser.Find(rune(valueobject.RootIdentifier)),
		},
	}

	r.BaseBlock.TreeNode.Val = r
	return r, nil
}
