package entity

import (
	"github.com/dddplayer/markdown/parser"
)

type Head struct {
	*BaseBlock
	Content string
	Level   int
}

func NewHead(p parser.Parser, l parser.Line) (*Head, error) {
	pr, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	n := NewNode()
	h := &Head{
		BaseBlock: &BaseBlock{
			node:   n,
			Parser: p,
		},
		Content: pr.Content(),
		Level:   len(pr.Identifier()),
	}

	n.B = h
	return h, err
}

func (h *Head) Continue(line parser.Line) ParseState {
	return Close
}
