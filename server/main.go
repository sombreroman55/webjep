package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
	api_router := r.PathPrefix("/api").Subrouter()
	api_router.HandleFunc("/clues", handleCluesPost).Methods("POST")
	api_router.HandleFunc("/host", handleHostGet).Methods("GET")
	r.Use(logMw)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
