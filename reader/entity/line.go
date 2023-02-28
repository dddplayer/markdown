package entity

type Line struct {
	Index     int
	FirstChar rune
	Content   []byte
}

func (l *Line) String() string {
	return string(l.Content)
}
