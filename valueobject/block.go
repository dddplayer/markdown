package valueobject

import (
	"github.com/dddplayer/markdown/datastructure"
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
)

type Block interface {
	IsOpen() bool
	Close() error
	Kind() valueobject.Kind
	Node() *datastructure.TreeNode
	AppendBlock(b Block)
	ParentBlock() Block
	Continue(line entity.Line) ParseDecision
}

type ParseDecision int

const (
	Continue ParseDecision = 1 << iota
	Children
	Close
)

type BlockState int

const (
	Opening BlockState = 1 << iota
	Closed
)

type BaseBlock struct {
	*datastructure.TreeNode
	state  BlockState
	Parser entity.Parser
}

func (b *BaseBlock) AppendBlock(block Block) {
	b.AppendChild(block.Node())
}

func (b *BaseBlock) IsOpen() bool {
	return b.state == Opening
}

func (b *BaseBlock) Close() error {
	b.state = Closed
	return nil
}

func (b *BaseBlock) Kind() valueobject.Kind {
	return b.Parser.Kind()
}

func (b *BaseBlock) Node() *datastructure.TreeNode {
	return b.TreeNode
}

func (b *BaseBlock) ParentBlock() Block {
	return b.Parent.Val.(Block)
}

func (b *BaseBlock) Continue(line entity.Line) ParseDecision {
	panic("should be override")
}
