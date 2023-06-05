package entity

import (
	"github.com/dddplayer/markdown/internal/domain/document/valueobject"
	"github.com/dddplayer/markdown/internal/domain/parser/entity"
	entity2 "github.com/dddplayer/markdown/pkg/datastructure"
)

type Head struct {
	*valueobject.BaseBlock
	Content string
	Level   int
}

func NewHead(p entity.BlockParser, l entity.Line) (*Head, error) {
	pr, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	h := &Head{
		BaseBlock: &valueobject.BaseBlock{
			TreeNode: entity2.EmptyTreeNode(),
			Parser:   p,
		},
		Content: pr.Content(),
		Level:   len(pr.Identifier()),
	}

	h.BaseBlock.TreeNode.Val = h
	return h, err
}

func (h *Head) Continue(line entity.Line) valueobject.ParseDecision {
	return valueobject.Close
}
