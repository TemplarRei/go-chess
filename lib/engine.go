package lib

import (
	"io"
)

type Board struct {
	fields  [8][8]field
	Turn    color
	Printer io.Writer
}

func (b *Board) String() {
	for _, row := range b.fields {
		for _, f := range row {
			b.Printer.Write([]byte(f.piece.String()))
		}
	}
}

type field struct {
	piece Figure
	color color
}

type Figure interface {
	String() string
	Moveset() []Move
}

// Field should have these?
// Shape   Shape
// Moveset Moveset
