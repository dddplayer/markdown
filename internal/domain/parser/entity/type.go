package entity

import (
	valueobject2 "github.com/dddplayer/markdown/internal/domain/parser/valueobject"
)

type Line interface {
	String() string
	FirstChar() rune
}

type ParseResult interface {
	Content() string
	Identifier() string
}

type BlockParser interface {
	Identifiers() []valueobject2.Identifier
	Kind() valueobject2.Kind
	Parse(l Line) (ParseResult, error)
}
