package entity

import (
	"github.com/dddplayer/markdown/internal/domain/parser/valueobject"
)

type Parser struct {
	pc *ParserCollection
}

func NewParser() *Parser {
	return &Parser{
		pc: newParserCollection(),
	}
}

func newParserCollection() *ParserCollection {
	pc := &ParserCollection{
		BlockParsers: []BlockParser{
			NewRoot(),
			NewHeading(),
			NewParagraph()},
		InlineParsers: nil,
		ParserMap:     map[valueobject.Identifier]BlockParser{},
	}
	pc.InitParserMap()

	return pc
}

func (p *Parser) Find(firstChar rune) BlockParser {
	parser := p.pc.ParserMap[valueobject.Identifier(firstChar)]
	if parser == nil {
		parser = p.pc.ParserMap[valueobject.EmptyIdentifier]
	}
	return parser
}
