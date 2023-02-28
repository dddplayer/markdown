package entity

import (
	"github.com/dddplayer/markdown/parser"
	"github.com/dddplayer/markdown/parser/valueobject"
)

type paragraph struct {
}

func NewParagraph() parser.Parser {
	return &paragraph{}
}

func (p *paragraph) Identifiers() []valueobject.Identifier {
	return nil
}

func (p *paragraph) Parse(l parser.Line) (parser.ParseResult, error) {
	return &block{
		identifier: "",
		content:    l.String(),
	}, nil
}

func (p *paragraph) Kind() valueobject.Kind {
	return valueobject.KindParagraph
}
