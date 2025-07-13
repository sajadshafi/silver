package game

import (
	"math"
	"silver/api/models"
)

func GetValidMoves(piece models.Piece, board models.Board, currentPosition int) []int {

	switch piece.PieceType {
	case models.Pawn:
		return GetValidPawnMoves(piece, board, currentPosition)
	case models.Rook:
		return GetValidRookMoves(piece, board, currentPosition)
	case models.Knight:
		return []int{}
	case models.Bishop:
		return GetValidBishopMoves(piece, board, currentPosition)
	case models.Queen:
		return []int{}
	case models.King:
		return []int{}
	default:
		return []int{}
	}
}

// func GetValidRookMoves(piece models.Piece, board models.Board, currentPosition int) []int {
// 	validMoves := []int{}
// 	return validMoves
// }

func GetValidRookMoves(piece models.Piece, board models.Board, currentPosition int) []int {
	validMoves := []int{}

	cell := GetFileRankFromIndex(currentPosition)
	start := currentPosition + 8
	for start <= 63 {
		validMoves = append(validMoves, start)
		start += 8
	}

	start = currentPosition - 8
	for start >= 0 {
		validMoves = append(validMoves, start)
		start -= 8
	}

	start = currentPosition + 1
	rankMin := 8 * (cell.Rank - 1)
	rankMax := rankMin + 7
	for start <= rankMax {
		validMoves = append(validMoves, start)
		start += 1
	}

	start = currentPosition - 1
	for start >= rankMin {
		validMoves = append(validMoves, start)
		start -= 1
	}

	return validMoves
}

func GetValidBishopMoves(piece models.Piece, board models.Board, currentPosition int) []int {
	validMoves := []int{}

	start := currentPosition + 9
	for start <= 63 {
		validMoves = append(validMoves, start)
		start += 9
	}

	start = currentPosition - 7
	for start >= 0 {
		validMoves = append(validMoves, start)
		start -= 7
	}

	start = currentPosition - 9
	for start >= 0 {
		validMoves = append(validMoves, start)
		start -= 9
	}

	start = currentPosition + 7
	for start <= 63 {
		validMoves = append(validMoves, start)
		start += 7
	}

	return validMoves
}

func GetValidPawnMoves(piece models.Piece, board models.Board, currentPosition int) []int {

	validMoves := []int{}
	// is it white or black - based on that whether we need to decrease the rank or increase the rank.
	cell := GetFileRankFromIndex(currentPosition)
	if piece.Color == models.White {
		possibleMove := currentPosition + 8
		if possibleMove <= 63 {
			validMoves = append(validMoves, possibleMove)
		}

		if cell.Rank == 2 {
			possibleSecondMove := possibleMove + 8
			if possibleSecondMove <= 63 {
				validMoves = append(validMoves, possibleSecondMove)
			}
		}
	} else {
		possibleMove := currentPosition - 8
		if possibleMove >= 0 {
			validMoves = append(validMoves, currentPosition-8)
		}

		if cell.Rank == 7 {
			possibleSecondMove := possibleMove - 8
			if possibleSecondMove >= 0 {
				validMoves = append(validMoves, possibleSecondMove)
			}
		}
	}

	// has the piece moved or not in the first place - so he can either move 1 steps or two steps.
	// is there a piece blocking its movement - if yes cannot move.
	// has a diagonal piece in front of it if yes is it oppsite color if yes then can capture also.

	return validMoves
}

func GetFileRankFromIndex(index int) models.Cell {
	file := (index % 8) + 1
	rank := int(math.Floor(float64(index)/8)) + 1

	return models.Cell{
		Rank: rank,
		File: file,
	}
}
