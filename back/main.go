package main

import (
	"encoding/json"
	"net/http"

	_ "github.com/google/uuid"
	_ "github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func handleNewGamePost(c echo.Context) error {
	log.Info(c.Request().Form)
	m := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&m)
	if err != nil {
		return err
	} else {
		for k, v := range m {
			log.Infof("%v: %v\n", k, v)
		}
	}

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

	e.POST("/newgame", handleNewGamePost)
	e.GET("/webjepws", handleWebjepWebsocketGet)

	e.Logger.Fatal(e.Start(":8000"))
}
