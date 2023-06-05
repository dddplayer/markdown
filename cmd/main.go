package main

import (
	"bytes"
	"fmt"
	"github.com/dddplayer/markdown/internal/application"
	entity2 "github.com/dddplayer/markdown/internal/domain/document/entity"
	mdvo "github.com/dddplayer/markdown/internal/domain/document/valueobject"
	psvo "github.com/dddplayer/markdown/internal/domain/parser/valueobject"
	"golang.org/x/tools/txtar"
	"os"
	"path/filepath"
)

func main() {

	dir, _ := os.MkdirTemp("", "hugo")
	defer os.RemoveAll(dir)

	var content = "# Clean Markdown\n" +
		"Markdown有着很好的写作体验..."
	md := "-- test.md --\n" + content
	writeFiles(md, dir)

	mdPath := filepath.Join(dir, "test.md")

	d, err := application.MDParse(mdPath)
	if err != nil {
		fmt.Println(err)
	}

	d.Step(
		func(block mdvo.Block) error {
			fmt.Println("in", block.Kind())
			switch block.Kind() {
			case psvo.KindRoot:
				fmt.Println("root")
			case psvo.KindHead:
				h := block.(*entity2.Head)
				fmt.Println(h.Level, h.Content)
			case psvo.KindParagraph:
				h := block.(*entity2.Paragraph)
				fmt.Println(h.Content)
			}
			return nil
		},
		func(block mdvo.Block) error {
			fmt.Println("out", block.Kind())
			return nil
		})
}

func writeFiles(s string, dir string) {
	data := txtar.Parse([]byte(s))

	for _, f := range data.Files {
		if err := os.WriteFile(
			filepath.Join(dir, f.Name),
			bytes.TrimSuffix(f.Data, []byte("\n")),
			os.ModePerm); err != nil {
			panic(err)
		}
	}
}
