package utils

import (
	"silver/engine/models"
	"strings"
	"unicode"
)

func ParseFEN(fen string) models.Board {
	rank := 7;
	file := 0;
	board := models.Board{};	
	var boardFen string;

	parts := strings.Split(fen, " ")
	if len(parts) > 0 {
		boardFen = parts[0];
	}

	for i := 0; i < len(boardFen); i++ {

		symbol := rune(boardFen[i]);

		if symbol == '/' {
			file = 0;
			rank -= 1;
		} else {
			if unicode.IsDigit(symbol) {
				file += int(symbol - '0');
			} else {
				var color models.Color
				
				if unicode.IsUpper(symbol) {
					color = models.White
				} else {
					color = models.Black
				}

				var pieceType models.PieceType = PieceTypeFromFen(unicode.ToLower(symbol));

				board[rank][file] = models.Square	{
					Piece: &models.Piece{
						PieceType: pieceType,
						Color:     color,
					},
				}

				file++;
			}
		}
	}

	return board;
}

func PieceTypeFromFen(symbol rune) models.PieceType {
	 var pieceMap = map[rune]models.PieceType {
		'k': models.King,
		'q': models.Queen,
		'r': models.Rook,
		'b': models.Bishop,
		'n': models.Knight,
		'p': models.Pawn,
	}

	return pieceMap[symbol];
}
