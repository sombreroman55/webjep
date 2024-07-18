package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)


func handleHomeGet(c echo.Context) error {
	// TODO: Serve SvelteKit SPA
	return nil
}

func handleLocalGamePost(c echo.Context) error {
	// TODO: Serve completely static version of the game page (no websocket)
	// TODO: Validate clues
	return nil
}

func handleGamePost(c echo.Context) error {
	// TODO: Create websocket server for WebRTC signalling and game synchronization
	// TODO: Validate clues
	return nil
}

func serveWebjep() {
	log.Info("Serving webjep")
	fmt.Println("Serving webjep!")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "site/static")
	e.GET("/", handleHomeGet)
	e.POST("/localgame", handleLocalGamePost)  // TBD: Do I need this?
	e.POST("/game", handleGamePost)

	e.Logger.Fatal(e.Start(":8000"))
}
