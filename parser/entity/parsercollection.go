package entity

import "github.com/dddplayer/markdown/parser/valueobject"

type ParserCollection struct {
	BlockParsers  []Parser
	InlineParsers []Parser
	ParserMap     map[valueobject.Identifier]Parser
}

func (m *ParserCollection) InitParserMap() {
	var ps []Parser
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
