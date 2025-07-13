package models

type Color string
type PieceType string

const (
	White Color = "white"
	Black Color = "black"

	Pawn   PieceType = "pawn"
	Rook   PieceType = "rook"
	Knight PieceType = "knight"
	Bishop PieceType = "bishop"
	Queen  PieceType = "queen"
	King   PieceType = "king"
)

type Piece struct {
	PieceType	PieceType `json:"type"`
	Color		Color     `json:"color"`
}

type Square struct {
	Piece *Piece `json:"piece,omitempty"`
}

type Board [8][8]Square

