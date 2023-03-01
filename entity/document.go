package entity

import (
	"fmt"
	"github.com/dddplayer/markdown/datastructure"
	"github.com/dddplayer/markdown/parser"
	"github.com/dddplayer/markdown/parser/valueobject"
	"github.com/dddplayer/markdown/reader"
	"github.com/dddplayer/markdown/reader/entity"
	"os"
)

type Document struct {
	*blockTree
	Name         string
	currentBlock Block
}

type StepIn func(block Block) error
type StepOut StepIn

func (d *Document) Step(in StepIn, out StepOut) {
	d.Walk(func(v any, ws datastructure.WalkState) datastructure.WalkStatus {
		b := v.(Block)
		if ws == datastructure.WalkIn {
			if err := in(b); err != nil {
				return datastructure.WalkStop
			}
		} else {
			if err := out(b); err != nil {
				return datastructure.WalkStop
			}
		}
		return datastructure.WalkContinue
	})
}

func (d *Document) Build(f *os.File) error {
	d.blockTree = NewTree()

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
	line := &line{l}
	p := parser.Find(line.FirstChar())
	switch p.Kind() {
	case valueobject.KindHead:
		return NewHead(p, line)
	case valueobject.KindParagraph:
		return NewParagraph(p, line)
	}

	return nil, nil
}
