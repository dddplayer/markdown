package valueobject

type Kind uint8

const (
	KindRoot Kind = iota
	KindHead
	KindParagraph
)
