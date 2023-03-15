package entity

import (
	"github.com/dddplayer/markdown/datastructure/entity"
	"github.com/dddplayer/markdown/parser"
	"github.com/dddplayer/markdown/parser/valueobject"
	valueobject2 "github.com/dddplayer/markdown/valueobject"
)

type Root struct {
	*valueobject2.BaseBlock
}

func NewRoot() (*Root, error) {
	r := &Root{
		BaseBlock: &valueobject2.BaseBlock{
			TreeNode: entity.EmptyTreeNode(),
			Parser:   parser.Find(rune(valueobject.RootIdentifier)),
		},
	}

	r.BaseBlock.TreeNode.Val = r
	return r, nil
}
