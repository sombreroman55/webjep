package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options
var connections = make(map[*websocket.Conn]bool)

func echo(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
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
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// TODO: Print the contents of the file to the console
func handleCluesPost(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, _, _ := r.FormFile("clues")
	defer file.Close()
	fmt.Println("Uploaded File contents:")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
	fmt.Println(r.Form)
}

func handleHostGet(w http.ResponseWriter, r *http.Request) {
	host := "Grimswald Bellingsly Himpleforthenworth"
	host_struct := struct {
		Host string `json:"host"`
	}{host}
	host_json, _ := json.Marshal(host_struct)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(host_json)
}

func main() {
	fmt.Println("Hello from WebJep!")

	r := mux.NewRouter()
	r.HandleFunc("/echo", echo)
	api_router := r.PathPrefix("/api").Subrouter()
	api_router.HandleFunc("/clues", handleCluesPost).Methods("POST")
	api_router.HandleFunc("/host", handleHostGet).Methods("GET")
	r.Use(logMw)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
