package board

import (
	"silver/api/models"
)

func InitializeBoard() models.Board {
	var board models.Board

	// Helper function to place a row of major pieces
	setupBackRank := func(color models.Color, start int) {
		board[start].Piece = &models.Piece{PieceType: models.Rook, Color: color}
		board[start+1].Piece = &models.Piece{PieceType: models.Knight, Color: color}
		board[start+2].Piece = &models.Piece{PieceType: models.Bishop, Color: color}
		board[start+3].Piece = &models.Piece{PieceType: models.Queen, Color: color}
		board[start+4].Piece = &models.Piece{PieceType: models.King, Color: color}
		board[start+5].Piece = &models.Piece{PieceType: models.Bishop, Color: color}
		board[start+6].Piece = &models.Piece{PieceType: models.Knight, Color: color}
		board[start+7].Piece = &models.Piece{PieceType: models.Rook, Color: color}
	}

	// Place black pieces
	setupBackRank(models.White, 0)
	for i := 8; i < 16; i++ {
		board[i].Piece = &models.Piece{PieceType: models.Pawn, Color: models.White}
	}

	// Place white pieces
	for i := 48; i < 56; i++ {
		board[i].Piece = &models.Piece{PieceType: models.Pawn, Color: models.Black}
	}
	setupBackRank(models.Black, 56)

	return board
}

func Piece(t, c string) *models.Piece {
	return &models.Piece{
		PieceType: models.PieceType(t),
		Color:     models.Color(c),
	}
}
