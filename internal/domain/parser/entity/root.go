package entity

import (
	valueobject2 "github.com/dddplayer/markdown/internal/domain/parser/valueobject"
)

type root struct {
}

func NewRoot() BlockParser {
	return &root{}
}

func (p *root) Identifiers() []valueobject2.Identifier {
	return []valueobject2.Identifier{valueobject2.RootIdentifier}
}

func (p *root) Parse(l Line) (ParseResult, error) {
	panic("root cannot parse anything")
}

func (p *root) Kind() valueobject2.Kind {
	return valueobject2.KindRoot
}
