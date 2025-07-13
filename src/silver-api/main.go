package main

import (
	"silver/api/game"
	"silver/api/models"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ValidMoveInput struct {
	Piece    models.Piece `json:"piece"`
	Board    models.Board `json:"board"`
	Position int          `json:"position"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default()) // All origins allowed by default
	router.GET("/game/initialize", func(c *gin.Context) {
		gameState := game.InitializeNewGame()
		c.JSON(200, gameState)
	})

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/game/get-valid-moves", func(c *gin.Context) {
		var json ValidMoveInput
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validMoves := game.GetValidMoves(json.Piece, json.Board, json.Position)
		c.JSON(200, validMoves)
	})

	router.Run("localhost:8080")
}
