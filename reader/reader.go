package reader

import (
	"github.com/dddplayer/markdown/reader/valueobject"
	"os"
)

func Scan(f *os.File, lineCB valueobject.ScanFunc) {
	lr := &valueobject.LineReader{F: f}
	lr.Scan(lineCB)
}
