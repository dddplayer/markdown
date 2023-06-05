package entity

import (
	valueobject2 "github.com/dddplayer/markdown/internal/domain/parser/valueobject"
)

type paragraph struct {
}

func NewParagraph() BlockParser {
	return &paragraph{}
}

func (p *paragraph) Identifiers() []valueobject2.Identifier {
	return nil
}

func (p *paragraph) Parse(l Line) (ParseResult, error) {
	return &block{
		identifier: "",
		content:    l.String(),
	}, nil
}

func (p *paragraph) Kind() valueobject2.Kind {
	return valueobject2.KindParagraph
}
