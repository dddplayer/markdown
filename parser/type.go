package parser

import "github.com/dddplayer/markdown/parser/valueobject"

type Line interface {
	String() string
	FirstChar() rune
}

type ParseResult interface {
	Content() string
	Identifier() string
}

type Parser interface {
	Identifiers() []valueobject.Identifier
	Kind() valueobject.Kind
	Parse(l Line) (ParseResult, error)
}
