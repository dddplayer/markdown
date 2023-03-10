package entity

import (
	"github.com/dddplayer/markdown/parser/valueobject"
)

type root struct {
}

func NewRoot() Parser {
	return &root{}
}

func (p *root) Identifiers() []valueobject.Identifier {
	return []valueobject.Identifier{valueobject.RootIdentifier}
}

func (p *root) Parse(l Line) (ParseResult, error) {
	panic("root cannot parse anything")
}

func (p *root) Kind() valueobject.Kind {
	return valueobject.KindRoot
}
