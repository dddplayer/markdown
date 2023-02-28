package entity

import (
	"github.com/dddplayer/markdown/parser"
	"strings"
)

type Paragraph struct {
	*BaseBlock
	Content []string
}

func NewParagraph(p parser.Parser, l parser.Line) (*Paragraph, error) {
	pr, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	n := NewNode()
	paragraph := &Paragraph{
		BaseBlock: &BaseBlock{
			node:   n,
			Parser: p,
		},
		Content: []string{pr.Content()},
	}

	n.B = paragraph
	return paragraph, err
}

func (p *Paragraph) Continue(line parser.Line) ParseState {
	pr, err := p.BaseBlock.Parser.Parse(line)
	if err != nil {
		panic("parse paragraph block err")
	}

	if isBlank(pr.Content()) {
		return Close
	}
	p.Content = append(p.Content, line.String())
	return Continue
}

func isBlank(s string) bool {
	tsl := strings.TrimSpace(s)
	if len(tsl) == 0 {
		return true
	}
	return false
}
