package entity

import (
	"bufio"
	"io"
	"os"
)

type ScanFunc func(l *Line) error

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
		err = f(NewLine(index, rawLine))
		if err != nil {
			break
		}
		index++
	}
}
