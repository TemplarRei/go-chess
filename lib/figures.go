package lib

type Shape string

const (
	WKing   Shape = "♔"
	WQueen  Shape = "♕"
	WRook   Shape = "♖"
	WBishop Shape = "♗"
	WKnight Shape = "♘"
	WPawn   Shape = "♙"
	WEmpty  Shape = "□"

	BKing   Shape = "♚"
	BQueen  Shape = "♛"
	BRook   Shape = "♜"
	BBishop Shape = "♝"
	BKnight Shape = "♞"
	BPawn   Shape = "♟"
	BEmpty  Shape = "■"
)

type color bool

const (
	White = iota
	Black = false
)

type Move struct{}
