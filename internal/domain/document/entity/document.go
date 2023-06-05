package entity

import (
	"fmt"
	"github.com/dddplayer/markdown/internal/domain/document/valueobject"
	parser "github.com/dddplayer/markdown/internal/domain/parser/entity"
	psvo "github.com/dddplayer/markdown/internal/domain/parser/valueobject"
	"github.com/dddplayer/markdown/internal/domain/reader/entity"
	entity2 "github.com/dddplayer/markdown/pkg/datastructure"
	"os"
)

type Document struct {
	*valueobject.BlockTree
	Name         string
	currentBlock valueobject.Block
	BlockParser  *parser.Parser
}

type StepFunc func(block valueobject.Block) error
type StepIn StepFunc
type StepOut StepFunc

func (d *Document) Step(in StepIn, out StepOut) {
	d.Walk(func(v any, ws entity2.WalkState) entity2.WalkStatus {
		b := v.(valueobject.Block)
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
	d.BlockTree = valueobject.NewTree(d.openRootBlock())

	parent := d.RootBlock()

	r := &entity.Reader{}
	r.Scan(f, func(l *entity.Line) error {
		if l == nil {
			panic("not support blank line in demo yet")
		}

	retry:
		if d.currentBlock != nil {
			state := d.currentBlock.Continue(NewLine(l))
			switch state {
			case valueobject.Children:
				panic("not implemented yet")
			case valueobject.Close:
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
			case valueobject.Continue:
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

func (d *Document) OpenBlock(l *entity.Line) (valueobject.Block, error) {
	line := NewLine(l)
	p := d.BlockParser.Find(line.FirstChar())
	switch p.Kind() {
	case psvo.KindHead:
		return NewHead(p, line)
	case psvo.KindParagraph:
		return NewParagraph(p, line)
	}

	return nil, nil
}

func (d *Document) openRootBlock() valueobject.Block {
	r, _ := NewRoot()
	return r
}
