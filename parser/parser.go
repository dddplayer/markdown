package parser

import (
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
)

func Find(firstChar rune) Parser {
	pc := &valueobject.ParserCollection{
		BlockParsers:  []Parser{entity.NewHeading(), entity.NewParagraph()},
		InlineParsers: nil,
		ParserMap:     map[valueobject.Identifier]Parser{},
	}
	pc.InitParserMap()

	p := pc.ParserMap[valueobject.Identifier(firstChar)]
	if p == nil {
		p = pc.ParserMap[valueobject.EmptyIdentifier]
	}
	return p
}
