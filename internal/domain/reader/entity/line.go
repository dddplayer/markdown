package entity

type Line struct {
	Index   int
	Content []byte
}

func (l *Line) String() string {
	return string(l.Content)
}

func NewLine(i int, b []byte) *Line {
	if len(b) == 0 {
		return nil
	}

	return &Line{
		Index:   i,
		Content: b,
	}
}
