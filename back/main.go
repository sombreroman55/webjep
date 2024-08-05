package main

// TODO: Create error pages and other routes
// FIXME: Validate form inputs!!!!!

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

// TBD: Replace this with a DB and a cache like Redis rather than in memory map like this?
var liveGames map[uuid.UUID]*Session

const MAX_PLAYERS = 5

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleNewGamePost(c echo.Context) error {
	log.Info("Got new game POST")
	var game Game
	err := json.NewDecoder(c.Request().Body).Decode(&game)
	if err != nil {
		log.Error(err)
		return err
	} else {
		log.Info(game)
	}

	gameId := uuid.New()
	hostId := uuid.New()
	session := newSession(gameId, game, hostId)
	liveGames[gameId] = session
	go session.start()

	cookie := new(http.Cookie)
	cookie.Name = "playerId"
	cookie.Value = hostId.String()
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	cookie = new(http.Cookie)
	cookie.Name = "isHost"
	cookie.Value = "true"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)

	gameUrl := fmt.Sprintf("/game?id=%s", gameId.String())
	return c.Redirect(http.StatusFound, gameUrl)
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
	session := liveGames[gameId]

	isHost, err := c.Cookie("isHost")
	var player *Player
	if err != nil {
		player = createNewPlayer(conn, session, false)
	} else {
		if isHost.Value == "true" {
			player = createNewPlayer(conn, session, true)
		}
	}

	if player.isHost {
		id, err := c.Cookie("isHost")
		if err != nil {
			player.id = uuid.MustParse(id.Value)
		} else {
			// TODO: This is bad, the host is corrupted
		}
	}

	player.session.register <- player

	go player.readPump()
	go player.writePump()
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
		gameId, err := uuid.Parse(gameIds)
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
		gameId, err := uuid.Parse(gameIds)
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
		if len(session.players) == MAX_PLAYERS {
			return nil
		}
		return next(c)
	}
}

func main() {
	log.Println("Serving Blockles site")
	log.SetLevel(log.InfoLevel)
	liveGames = make(map[uuid.UUID]*Session)

	e := echo.New()
	e.Use(middleware.CORS())

	// TODO: Put SvelteKit-generated files here
	e.GET("/", handleGameGet, gameMiddleware)
	e.GET("/game-builder", handleGameGet, gameMiddleware)
	e.GET("/local-game", handleGameGet, gameMiddleware)
	e.GET("/new-game", handleGameGet, gameMiddleware)
	e.GET("/game", handleGameGet, gameMiddleware)

	e.POST("/new-game", handleNewGamePost)
	e.GET("/gamews", handleWebjepWebsocketGet, gameMiddleware)

	e.Logger.Fatal(e.Start(":8000"))
}
