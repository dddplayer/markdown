package entity

import (
	"os"
)

type Reader struct {
}

func (r *Reader) Scan(f *os.File, lineCB ScanFunc) {
	lr := &LineReader{F: f}
	lr.Scan(lineCB)
}
