package valueobject

import (
	"bufio"
	"github.com/dddplayer/markdown/reader/entity"
	"github.com/dddplayer/markdown/reader/factory"
	"io"
	"os"
)

type ScanFunc func(l *entity.Line) error

type LineReader struct {
	F *os.File
}

func (l *LineReader) Scan(f ScanFunc) {
	reader := bufio.NewReader(l.F)
	index := 1
	for {
		rawLine, isPrefix, err := reader.ReadLine()
		if isPrefix {
			panic("line is too long")
		}
		if err == io.EOF {
			break
		}
		err = f(factory.NewLine(index, rawLine))
		if err != nil {
			break
		}
		index++
	}
}
