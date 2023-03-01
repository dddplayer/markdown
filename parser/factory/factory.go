package factory

import (
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
)

func NewParserCollection() *entity.ParserCollection {
	pc := &entity.ParserCollection{
		BlockParsers: []entity.Parser{
			entity.NewRoot(),
			entity.NewHeading(),
			entity.NewParagraph()},
		InlineParsers: nil,
		ParserMap:     map[valueobject.Identifier]entity.Parser{},
	}
	pc.InitParserMap()

	return pc
}
