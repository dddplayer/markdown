package entity

import (
	"github.com/dddplayer/markdown/parser"
	"github.com/dddplayer/markdown/parser/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
)

type Root struct {
	*BaseBlock
}

func NewRoot() (*Root, error) {
	n := NewNode()
	r := &Root{
		BaseBlock: &BaseBlock{
			blockNode: n,
			Parser:    parser.Find(rune(valueobject.RootIdentifier)),
		},
	}

	n.MdBlock = r
	return r, nil
}

func (r *Root) Continue(line entity.Line) ParseState {
	return Close
}
