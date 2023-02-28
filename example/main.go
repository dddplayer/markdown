package main

import (
	"bytes"
	"fmt"
	"github.com/dddplayer/markdown"
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
	fmt.Printf("%#v", d)
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
