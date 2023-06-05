package entity

import (
	"github.com/dddplayer/markdown/internal/domain/reader/entity"
)

type line struct {
	*entity.Line
}

func (l *line) FirstChar() rune {
	return rune(l.Line.Content[0])
}

func NewLine(l *entity.Line) *line {
	return &line{l}
}
