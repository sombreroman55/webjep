package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ServerError struct {
  Message string
}

func (e *ServerError) Error() string {
  return e.Message
}

var jepUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:5173"
	},
}

var rtcUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:5173"
	},
}

var connections = make(map[*websocket.Conn]bool)
var activeGames = make(map[string]GameState)

func jepSocket(w http.ResponseWriter, r *http.Request) {
	c, err := jepUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	connections[c] = true

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		for conn := range connections {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func rtcSocket(w http.ResponseWriter, r *http.Request) {
	c, err := rtcUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	connections[c] = true

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		for conn := range connections {
			err = conn.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func logMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func handleCluesPost(w http.ResponseWriter, r *http.Request) {
	game := CreateNewGame(r)
	activeGames[game.Id] = game
	w.Write([]byte(`{"status": "ok", "gameId": "` + game.Id + `"}`))
	fmt.Println(r.Form)
}

func handleHostGet(w http.ResponseWriter, r *http.Request) {
	host := "Grimswald Bellingsly Himpleforthenworth"
	host_struct := struct {
		Host string `json:"host"`
	}{host}
	host_json, _ := json.Marshal(host_struct)
	w.Write(host_json)
}

func main() {
	fmt.Println("Hello from WebJep!")

	r := mux.NewRouter()
	r.HandleFunc("/jep", jepSocket)
	r.HandleFunc("/rtc", rtcSocket)
	api_router := r.PathPrefix("/api").Subrouter()
	api_router.HandleFunc("/clues", handleCluesPost).Methods("POST")
	api_router.HandleFunc("/host", handleHostGet).Methods("GET")
	r.Use(logMw)

	corsOptions := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsHandler := handlers.CORS(corsOptions)

	http.Handle("/", corsHandler(r))
	http.ListenAndServe(":8080", nil)
}
