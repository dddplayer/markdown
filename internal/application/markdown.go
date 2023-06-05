package application

import (
	"errors"
	"github.com/dddplayer/markdown/internal/domain/document/entity"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func MDParse(path string) (*entity.Document, error) {
	filename := filepath.Base(path)

	ext := filepath.Ext(filename)
	if ext != ".md" {
		check(errors.New("only .md file supported"))
	}

	f, err := os.Open(path)
	defer f.Close()

	check(err)

	return Parse(filename, f)
}
