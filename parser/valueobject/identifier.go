package valueobject

type Identifier rune

const (
	EmptyIdentifier Identifier = Identifier(rune(0))
	HeadIdentifier  Identifier = Identifier(rune('#'))
)
