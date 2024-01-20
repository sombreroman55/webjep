package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type GameError struct {
  Message string
}

func (e *GameError) Error() string {
  return e.Message
}

type Host struct {
	Name   string
	Token  string
	Socket *websocket.Conn
}

type Player struct {
	Name   string
	Wager  int
	Score  int
	Socket *websocket.Conn
}

type Clue struct {
	Question string
	Answer   string
	Value    int
}

type Category struct {
	Title string
	Clues [5]Clue
}

type Round struct {
	Categories [6]Category
}

type FinalRound struct {
	Category string
	Question string
	Answer   string
}

type Game struct {
	SingleJep Round
	DoubleJep Round
	FinalJep  FinalRound
}

type GameState struct {
	Game            Game
	Id              string
	State           string
	Host            Host
	Players         []Player
	LastPlayer      Player
	CurrentRound    int
	CurrentCategory Category
	CurrentClue     Clue
}

func parseClues(file []byte) Game {
	var game Game
	json.Unmarshal(file, &game)
	return game
}

func CreateNewGame(r *http.Request) GameState {
	log.Println("Creating new game...")
	var gameState GameState
	file, _, _ := r.FormFile("clues")
	hostName := r.FormValue("host")
	defer file.Close()
	var fileContents []byte
	file.Read(fileContents)
	gameState.Game = parseClues(fileContents)
	gameState.Id = uuid.New().String()
	gameState.State = "pending"
	gameState.Host = Host{
		Name:  hostName,
		Token: uuid.New().String(),
	}
	gameState.Players = make([]Player, 0)
	gameState.CurrentRound = 1
	return gameState
}

func (gs *GameState) AddPlayer(r *http.Request) {
	var player Player
	player.Name = r.FormValue("name")
	if len(gs.Players) == 0 {
		gs.LastPlayer = player
	}
	gs.Players = append(gs.Players, player)
}

func (gs *GameState) StartGame() {
	gs.State = "in_progress"
	gs.Host.Socket.WriteMessage(websocket.TextMessage, []byte("Game started!"))
	for _, player := range gs.Players {
		player.Socket.WriteMessage(websocket.TextMessage, []byte("Game started!"))
	}
}
