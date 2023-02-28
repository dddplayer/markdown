package markdown

import (
	"errors"
	"github.com/dddplayer/markdown/entity"
	"github.com/dddplayer/markdown/service"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Parse(path string) (*entity.Document, error) {
	filename := filepath.Base(path)

	ext := filepath.Ext(filename)
	if ext != ".md" {
		check(errors.New("only .md file supported"))
	}

	f, err := os.Open(path)
	defer f.Close()

	check(err)

	return service.Parse(filename, f)
}
