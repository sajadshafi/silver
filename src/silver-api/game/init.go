package game

import (
	"silver/api/board"
	"silver/api/models"
)

func InitializeNewGame() models.Game {

	var board models.Board = board.InitializeBoard()
	return models.Game{
		GameId: "",
		Board:  board,
	}
} 
