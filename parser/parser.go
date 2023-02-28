package parser

import (
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
)

func Find(firstChar rune) entity.Parser {
	pc := &entity.ParserCollection{
		BlockParsers: []entity.Parser{
			entity.NewRoot(),
			entity.NewHeading(),
			entity.NewParagraph()},
		InlineParsers: nil,
		ParserMap:     map[valueobject.Identifier]entity.Parser{},
	}
	pc.InitParserMap()

	p := pc.ParserMap[valueobject.Identifier(firstChar)]
	if p == nil {
		p = pc.ParserMap[valueobject.EmptyIdentifier]
	}
	return p
}
