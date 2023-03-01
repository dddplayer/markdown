package factory

import "github.com/dddplayer/markdown/reader/entity"

func NewLine(i int, b []byte) *entity.Line {
	if len(b) == 0 {
		return nil
	}

	return &entity.Line{
		Index:   i,
		Content: b,
	}
}
