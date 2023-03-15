package entity

import (
	"fmt"
	entity2 "github.com/dddplayer/markdown/datastructure/entity"
	"github.com/dddplayer/markdown/parser"
	psvo "github.com/dddplayer/markdown/parser/valueobject"
	"github.com/dddplayer/markdown/reader"
	"github.com/dddplayer/markdown/reader/entity"
	mdvo "github.com/dddplayer/markdown/valueobject"
	"os"
)

type Document struct {
	*mdvo.BlockTree
	Name         string
	currentBlock mdvo.Block
}

type StepIn func(block mdvo.Block) error
type StepOut StepIn

func (d *Document) Step(in StepIn, out StepOut) {
	d.Walk(func(v any, ws entity2.WalkState) entity2.WalkStatus {
		b := v.(mdvo.Block)
		if ws == entity2.WalkIn {
			if err := in(b); err != nil {
				return entity2.WalkStop
			}
		} else {
			if err := out(b); err != nil {
				return entity2.WalkStop
			}
		}
		return entity2.WalkContinue
	})
}

func (d *Document) Build(f *os.File) error {
	d.BlockTree = mdvo.NewTree(d.openRootBlock())

	parent := d.RootBlock()
	reader.Scan(f, func(l *entity.Line) error {
		if l == nil {
			panic("not support blank line in demo yet")
		}

	retry:
		if d.currentBlock != nil {
			state := d.currentBlock.Continue(&line{l})
			switch state {
			case mdvo.Children:
				panic("not implemented yet")
			case mdvo.Close:
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
			case mdvo.Continue:
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

func (d *Document) OpenBlock(l *entity.Line) (mdvo.Block, error) {
	line := &line{l}
	p := parser.Find(line.FirstChar())
	switch p.Kind() {
	case psvo.KindHead:
		return NewHead(p, line)
	case psvo.KindParagraph:
		return NewParagraph(p, line)
	}

	return nil, nil
}

func (d *Document) openRootBlock() mdvo.Block {
	r, _ := NewRoot()
	return r
}
