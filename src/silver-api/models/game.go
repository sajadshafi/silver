package models;

type Game struct {
	GameId  string `json:"gameId"`
	Board	Board `json:"board"`
}

