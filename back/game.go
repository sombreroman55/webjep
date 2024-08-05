package main

import (
	"strings"
	"strconv"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Clue struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Value    int
	Answered bool
}

type Category struct {
	Title string `json:"title"`
	Clues []Clue `json:"clues"`
}

type Round struct {
	Categories []Category `json:"categories"`
}

type FinalRound struct {
	Category string `json:"category"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type GameData struct {
	Rounds     []Round     `json:"rounds"`
	FinalRound *FinalRound `json:"finalRound"`
}

type GameSettings struct {
	PlayFinalJep         bool    `json:"playFinalJep"`
	BaseClueValue        int     `json:"baseClueValue"`
	InterClueInterval    float64 `json:"interClueInterval"`
	InterRoundMultiplier float64 `json:"InterRoundMultiplier"`
	NumRounds            int     `json:"numberOfRounds"`
	CategoriesPerRound   int     `json:"categoriesPerRound"`
	CluesPerCategory     int     `json:"cluesPerCategory"`
}

type Game struct {
	Data     GameData     `json:"data"`
	Settings GameSettings `json:"settings"`
}

type GameState int

const (
	ready = iota
	locked
	inProgress
	finished
	cancelled
)

type ClueIndices struct {
	r int
	c int
	q int
}

type Session struct {
	id            uuid.UUID
	game          Game
	host          uuid.UUID
	state         int
	currentClue   ClueIndices
	currentPlayer uuid.UUID
	players       map[uuid.UUID]*Player
	commands      chan []byte
	register      chan *Player
	unregister    chan *Player
}

func computeClueValues(game *Game) {
	for r := range game.Data.Rounds {
		for c := range game.Data.Rounds[r].Categories {
			for q := range game.Data.Rounds[r].Categories[c].Clues {
				bcv := game.Settings.BaseClueValue
				irm := game.Settings.InterRoundMultiplier
				ici := game.Settings.InterClueInterval
				crbv := int(float64(bcv) + float64(r)*float64(bcv)*irm)
				ccv := int(float64(crbv) + float64(q)*float64(crbv)*ici)
				game.Data.Rounds[r].Categories[c].Clues[q].Value = ccv
				game.Data.Rounds[r].Categories[c].Clues[q].Answered = false
			}
		}
	}
}

func newSession(id uuid.UUID, game Game, host uuid.UUID) *Session {
	return &Session{
		id:         id,
		game:       game,
		host:       host,
		state:      ready,
		players:    make(map[uuid.UUID]*Player),
		commands:   make(chan []byte),
		register:   make(chan *Player),
		unregister: make(chan *Player),
	}
}

func (s *Session) start() {
	for {
		select {
		case player := <-s.register:
			s.players[player.id] = player
		case player := <-s.unregister:
			if _, ok := s.players[player.id]; ok {
				delete(s.players, player.id)
				close(player.send)
			}
		case commandBytes := <-s.commands:
			// TODO: Update game state here...
			// FIXME: Validate this input!!!!!
			commSlice := strings.Split(string(commandBytes), ",")
			command := commSlice[0]
			log.Info(command)
			switch command {
			case "buzz":
				s.currentPlayer = uuid.MustParse(commSlice[1])
				player := s.players[s.currentPlayer]
				player.canBuzz = false
			case "clue":
				s.currentClue.r, _ = strconv.Atoi(commSlice[1])
				s.currentClue.c, _ = strconv.Atoi(commSlice[2])
				s.currentClue.q, _ = strconv.Atoi(commSlice[3])
			case "correct":
				player := s.players[s.currentPlayer]
				r, c, q := s.currentClue.r, s.currentClue.c, s.currentClue.q
				clue := s.game.Data.Rounds[r].Categories[c].Clues[q]
				player.score += clue.Value
			case "incorrect":
				player := s.players[s.currentPlayer]
				r, c, q := s.currentClue.r, s.currentClue.c, s.currentClue.q
				clue := s.game.Data.Rounds[r].Categories[c].Clues[q]
				player.score -= clue.Value
			}
			// TODO: ...then send the updated state back to clients (or at least the deltas)
		}
	}
}

func (g *Game) setCurrentClue(c Clue) {

}
