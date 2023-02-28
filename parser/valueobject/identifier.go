package valueobject

type Identifier rune

const (
	RootIdentifier  Identifier = Identifier(rune(-1))
	EmptyIdentifier Identifier = Identifier(rune(0))
	HeadIdentifier  Identifier = Identifier(rune('#'))
)
