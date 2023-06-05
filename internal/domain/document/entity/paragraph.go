package entity

import (
	"github.com/dddplayer/markdown/internal/domain/document/valueobject"
	"github.com/dddplayer/markdown/internal/domain/parser/entity"
	entity2 "github.com/dddplayer/markdown/pkg/datastructure"
	"strings"
)

type Paragraph struct {
	*valueobject.BaseBlock
	Content []string
}

func NewParagraph(p entity.BlockParser, l entity.Line) (*Paragraph, error) {
	pr, err := p.Parse(l)
	if err != nil {
		return nil, err
	}

	paragraph := &Paragraph{
		BaseBlock: &valueobject.BaseBlock{
			TreeNode: entity2.EmptyTreeNode(),
			Parser:   p,
		},
		Content: []string{pr.Content()},
	}

	paragraph.BaseBlock.TreeNode.Val = paragraph
	return paragraph, err
}

func (p *Paragraph) Continue(line entity.Line) valueobject.ParseDecision {
	pr, err := p.BaseBlock.Parser.Parse(line)
	if err != nil {
		panic("parse paragraph block err")
	}

	if isBlank(pr.Content()) {
		return valueobject.Close
	}
	p.Content = append(p.Content, line.String())
	return valueobject.Continue
}

func isBlank(s string) bool {
	tsl := strings.TrimSpace(s)
	if len(tsl) == 0 {
		return true
	}
	return false
}
