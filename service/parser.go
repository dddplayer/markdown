package service

import (
	"github.com/dddplayer/markdown/entity"
	"os"
)

func Parse(name string, f *os.File) (*entity.Document, error) {
	d := &entity.Document{Name: name}
	if err := d.Build(f); err != nil {
		return nil, err
	}

	return d, nil
}
