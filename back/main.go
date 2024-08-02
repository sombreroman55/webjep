package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type ServerSession struct {
	id       uuid.UUID
	gameData Session
}

func newServerSession(id uuid.UUID, data Session) *ServerSession {
	return &ServerSession{
		id: id,
		gameData: data,
	}
}

var liveGames map[uuid.UUID]*ServerSession

func handleNewGamePost(c echo.Context) error {
	log.Info("Got new game POST")
	var session Session
	err := json.NewDecoder(c.Request().Body).Decode(&session)
	if err != nil {
		log.Error(err)
		return err
	} else {
		log.Info(session)
	}

	gameId := uuid.New()
	servSession := newServerSession(gameId, session)

	// TODO: Set user cookies
	// TODO: Redirect user to game page with game ID
	// TODO: Tell them to distribute the link on the front end

	return c.String(http.StatusOK, "Thanks for the form!")
}

func handleWebjepWebsocketGet(c echo.Context) error {
	return c.String(http.StatusOK, "This will connect you to a websocket")
}

// TODO: Eventually, serve the entire application from this binary as well
func main() {
	log.Println("Serving Blockles site")
	log.SetLevel(log.InfoLevel)

	e := echo.New()
	e.Use(middleware.CORS())

	e.POST("/api/new-game", handleNewGamePost)
	e.POST("/api/webjepws", handleWebjepWebsocketGet)

	e.Logger.Fatal(e.Start(":8000"))
}
