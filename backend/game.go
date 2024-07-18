package main

import (
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type SignalServer struct {
}

type GameRoom struct {
	signaller SignalServer
	players   []Player
	game      Game
}
