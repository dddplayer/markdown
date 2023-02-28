package main

import (
	"bytes"
	"fmt"
	"github.com/dddplayer/markdown"
	"github.com/dddplayer/markdown/entity"
	"github.com/dddplayer/markdown/parser/valueobject"
	"os"
	"path/filepath"

	"golang.org/x/tools/txtar"
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
		func(block entity.Block) error {
			fmt.Println("in", block.Kind())
			switch block.Kind() {
			case valueobject.KindRoot:
				fmt.Println("root")
			case valueobject.KindHead:
				h := block.(*entity.Head)
				fmt.Println(h.Level, h.Content)
			case valueobject.KindParagraph:
				h := block.(*entity.Paragraph)
				fmt.Println(h.Content)
			}
			return nil
		},
		func(block entity.Block) error {
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
