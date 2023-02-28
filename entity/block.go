package entity

import (
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
)

type Block interface {
	IsOpen() bool
	Close() error
	Kind() valueobject.Kind
	Node() *node
	AppendBlock(b Block)
	ParentBlock() Block
	Continue(line entity.Line) ParseState
}

type ParseState int

const (
	Continue ParseState = 1 << iota
	Children
	Close
)

type BlockState int

const (
	Opening BlockState = 1 << iota
	Closed
)

type BaseBlock struct {
	*node
	state  BlockState
	Parser entity.Parser
}

func (b *BaseBlock) AppendBlock(block Block) {
	b.node.AppendChild(block.Node().Node)
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

func (b *BaseBlock) Node() *node {
	return b.node
}

func (b *BaseBlock) ParentBlock() Block {
	return b.node.Parent.Val.(*node).B
}

func (b *BaseBlock) Continue(line entity.Line) ParseState {
	panic("should be override")
}
