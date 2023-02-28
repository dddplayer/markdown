package entity

import (
	"fmt"
	"github.com/dddplayer/markdown/parser"
	"github.com/dddplayer/markdown/parser/valueobject"
	"github.com/dddplayer/markdown/reader"
	"github.com/dddplayer/markdown/reader/entity"
	"os"
)

type Document struct {
	*tree
	Name         string
	currentBlock Block
}

func (d *Document) Build(f *os.File) error {
	d.tree = NewTree()

	parent := d.RootBlock()
	reader.Scan(f, func(l *entity.Line) error {
		if l == nil {
			panic("not support blank line in demo yet")
		}

	retry:
		if d.currentBlock != nil {
			state := d.currentBlock.Continue(&line{l})
			switch state {
			case Children:
				panic("not implemented yet")
			case Close:
				if err := d.currentBlock.Close(); err != nil {
					return err
				}
				if parent.IsOpen() {
					d.currentBlock = parent
					parent = parent.ParentBlock()
				} else {
					d.currentBlock = nil
				}
				goto retry
			case Continue:
				fmt.Println("continue")
			}
		}
		ob, err := d.OpenBlock(l)
		if err != nil {
			return err
		}
		d.currentBlock = ob
		parent.AppendBlock(ob)
		return nil
	})
	return nil
}

func (d *Document) OpenBlock(l *entity.Line) (Block, error) {
	p := parser.Find(l.FirstChar)
	switch p.Kind() {
	case valueobject.KindHead:
		return NewHead(p, &line{l})
	case valueobject.KindParagraph:
		return NewParagraph(p, &line{l})
	}

	return nil, nil
}