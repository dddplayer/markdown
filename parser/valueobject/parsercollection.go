package valueobject

import (
	"github.com/dddplayer/markdown/parser"
)

type ParserCollection struct {
	BlockParsers  []parser.Parser
	InlineParsers []parser.Parser
	ParserMap     map[Identifier]parser.Parser
}

func (m *ParserCollection) InitParserMap() {
	var ps []parser.Parser
	ps = append(ps, m.BlockParsers...)
	ps = append(ps, m.InlineParsers...)

	for _, p := range ps {
		is := p.Identifiers()
		if is == nil {
			m.ParserMap[EmptyIdentifier] = p
		} else {
			for _, i := range is {
				m.ParserMap[i] = p
			}
		}
	}
}
