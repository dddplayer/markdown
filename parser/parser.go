package parser

import (
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/factory"
	"github.com/dddplayer/markdown/parser/valueobject"
)

func Find(firstChar rune) entity.Parser {
	pc := factory.NewParserCollection()

	p := pc.ParserMap[valueobject.Identifier(firstChar)]
	if p == nil {
		p = pc.ParserMap[valueobject.EmptyIdentifier]
	}
	return p
}
