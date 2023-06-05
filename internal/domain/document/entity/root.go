package entity

import (
	valueobject2 "github.com/dddplayer/markdown/internal/domain/document/valueobject"
	"github.com/dddplayer/markdown/internal/domain/parser/entity"
	"github.com/dddplayer/markdown/internal/domain/parser/valueobject"
	"github.com/dddplayer/markdown/pkg/datastructure"
)

type Root struct {
	*valueobject2.BaseBlock
}

func NewRoot() (*Root, error) {
	p := entity.NewParser()
	r := &Root{
		BaseBlock: &valueobject2.BaseBlock{
			TreeNode: datastructure.EmptyTreeNode(),
			Parser:   p.Find(rune(valueobject.RootIdentifier)),
		},
	}

	r.BaseBlock.TreeNode.Val = r
	return r, nil
}
