package entity

import (
	"github.com/dddplayer/markdown/datastructure"
	"github.com/dddplayer/markdown/parser/entity"
)

type Head struct {
	*BaseBlock
	Content string
	Level   int
}

func NewHead(p entity.Parser, l entity.Line) (*Head, error) {
	pr, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	h := &Head{
		BaseBlock: &BaseBlock{
			TreeNode: datastructure.EmptyTreeNode(),
			Parser:   p,
		},
		Content: pr.Content(),
		Level:   len(pr.Identifier()),
	}

	h.BaseBlock.TreeNode.Val = h
	return h, err
}

func (h *Head) Continue(line entity.Line) ParseState {
	return Close
}
