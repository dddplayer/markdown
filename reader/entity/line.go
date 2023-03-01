package entity

type Line struct {
	Index   int
	Content []byte
}

func (l *Line) String() string {
	return string(l.Content)
}
