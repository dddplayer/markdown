package entity

import "github.com/dddplayer/markdown/reader/entity"

type line struct {
	*entity.Line
}

func (l *line) FirstChar() rune {
	return rune(l.Line.Content[0])
}
