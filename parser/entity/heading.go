package entity

import (
	"errors"
	"github.com/dddplayer/markdown/parser/valueobject"
	"golang.org/x/exp/slices"
	"strings"
)

type heading struct {
}

func NewHeading() Parser {
	return &heading{}
}

func (h *heading) Kind() valueobject.Kind {
	return valueobject.KindHead
}

func (h *heading) Identifiers() []valueobject.Identifier {
	return []valueobject.Identifier{valueobject.HeadIdentifier}
}

func (h *heading) Parse(l Line) (ParseResult, error) {
	if h.valid(l) == false {
		return nil, errors.New("invalid line for heading block parser")
	}

	return &block{
		identifier: level(l.String()),
		content:    headerRaw(l.String()),
	}, nil
}

func (h *heading) valid(l Line) bool {
	return slices.Contains(h.Identifiers(), valueobject.Identifier(l.FirstChar()))
}

func headerRaw(s string) string {
	return strings.TrimSpace(strings.TrimLeft(s, "#"))
}

func level(s string) string {
	level := len(s) - len(strings.TrimLeft(s, "#"))
	return s[:level]
}
