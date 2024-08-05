package main

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// TODO: Also add WebRTC signalling data here
type Player struct {
	name   string
	score  int
	conn   *websocket.Conn
	isHost bool
	id     uuid.UUID
}

func createNewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		name:   "",
		score:  0,
		conn:   conn,
		isHost: false,
		id:     uuid.New(),
	}
}
