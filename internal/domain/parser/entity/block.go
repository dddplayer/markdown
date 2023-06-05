package entity

type block struct {
	identifier string
	content    string
}

func (b *block) Content() string {
	return b.content
}
func (b *block) Identifier() string {
	return b.identifier
}
