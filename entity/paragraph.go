package entity

import (
	"github.com/dddplayer/markdown/datastructure"
	"github.com/dddplayer/markdown/parser/entity"
	"strings"
)

type Paragraph struct {
	*BaseBlock
	Content []string
}

func NewParagraph(p entity.Parser, l entity.Line) (*Paragraph, error) {
	pr, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	paragraph := &Paragraph{
		BaseBlock: &BaseBlock{
			TreeNode: datastructure.EmptyTreeNode(),
			Parser:   p,
		},
		Content: []string{pr.Content()},
	}

	paragraph.BaseBlock.TreeNode.Val = paragraph
	return paragraph, err
}

func (p *Paragraph) Continue(line entity.Line) ParseDecision {
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
