package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type GameState int

const (
	ready = iota
	locked
	inProgress
	finished
	cancelled
)

type ServerSession struct {
	id       uuid.UUID
	gameData Session
	host     uuid.UUID
	state    int
	players  []*Player
}

func newServerSession(id uuid.UUID, data Session) *ServerSession {
	return &ServerSession{
		id:       id,
		gameData: data,
	}
}

// TBD: Replace this with a DB and a cache like Redis rather than in memory map like this
var liveGames map[uuid.UUID]*ServerSession

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

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
	liveGames[gameId] = servSession
	// TODO: Kick off session goroutine here
	// TODO: Set user cookies
	// TODO: Redirect user to game page with game ID
	// TODO: Tell them to distribute the link on the front end

	return c.String(http.StatusOK, "Thanks for the form!")
}

func handleWebjepWebsocketGet(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	// Info was validated in middleware, just use it directly here
	gameIds := c.QueryParam("id")
	gameId := uuid.MustParse(gameIds)
	player := createNewPlayer(conn)
	session := liveGames[gameId]
	session.gameData.addPlayer(player)

	// TODO: Kick off player goroutines here
	return nil
}

func handleGameGet(c echo.Context) error {
	// TODO: Serve the svelte page here
	return nil
}

func websocketMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		gameIds := c.QueryParam("id")
		if gameIds == "" {
			return nil
		}
		gameId,err := uuid.Parse(gameIds)
		if err != nil {
			return nil
		}
		session, ok := liveGames[gameId]
		if !ok {
			return nil
		}
		if session.state != inProgress {
			return nil
		}
		return next(c)
	}
}

func gameMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		gameIds := c.QueryParam("id")
		if gameIds == "" {
			return nil
		}
		gameId,err := uuid.Parse(gameIds)
		if err != nil {
			return nil
		}
		session, ok := liveGames[gameId]
		if !ok {
			return nil
		}
		if session.state != ready {
			return nil
		}
		if len(session.players) == 5 {
			return nil
		}
		return next(c)
	}
}

// TODO: Eventually, serve the entire application from this binary as well
func main() {
	log.Println("Serving Blockles site")
	log.SetLevel(log.InfoLevel)
	liveGames = make(map[uuid.UUID]*ServerSession)

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/game", handleGameGet, gameMiddleware)
	e.POST("/api/new-game", handleNewGamePost)
	e.POST("/api/webjepws", handleWebjepWebsocketGet, gameMiddleware)

	e.Logger.Fatal(e.Start(":8000"))
}
