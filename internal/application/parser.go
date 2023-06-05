package application

import (
	"github.com/dddplayer/markdown/internal/domain/document/entity"
	entity2 "github.com/dddplayer/markdown/internal/domain/parser/entity"
	"os"
)

func Parse(name string, f *os.File) (*entity.Document, error) {
	d := &entity.Document{
		Name:        name,
		BlockParser: entity2.NewParser(),
	}
	if err := d.Build(f); err != nil {
		return nil, err
	}

	return d, nil
}
