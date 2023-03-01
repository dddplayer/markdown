package main

import (
	"bytes"
	"fmt"
	"github.com/dddplayer/markdown"
	"github.com/dddplayer/markdown/entity"
	psvo "github.com/dddplayer/markdown/parser/valueobject"
	mdvo "github.com/dddplayer/markdown/valueobject"
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

	d, err := markdown.Parse(mdPath)
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
				h := block.(*entity.Head)
				fmt.Println(h.Level, h.Content)
			case psvo.KindParagraph:
				h := block.(*entity.Paragraph)
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
