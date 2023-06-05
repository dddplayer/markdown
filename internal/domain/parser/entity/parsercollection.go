package entity

import (
	"github.com/dddplayer/markdown/internal/domain/parser/valueobject"
)

type ParserCollection struct {
	BlockParsers  []BlockParser
	InlineParsers []BlockParser
	ParserMap     map[valueobject.Identifier]BlockParser
}

func (m *ParserCollection) InitParserMap() {
	var ps []BlockParser
	ps = append(ps, m.BlockParsers...)
	ps = append(ps, m.InlineParsers...)

	for _, p := range ps {
		is := p.Identifiers()
		if is == nil {
			m.ParserMap[valueobject.EmptyIdentifier] = p
		} else {
			for _, i := range is {
				m.ParserMap[i] = p
			}
		}
	}
}
