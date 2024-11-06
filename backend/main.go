package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type Message struct {
	Text string `json:"text"`
}

var messages []Message
var mutex = &sync.Mutex{}

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	json.NewEncoder(w).Encode(messages)
	mutex.Unlock()
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	json.NewDecoder(r.Body).Decode(&message)
	mutex.Lock()
	messages = append(messages, message)
	mutex.Unlock()
	w.WriteHeader(http.StatusCreated)
}

func main() {
	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postMessage(w, r)
		} else {
			getMessages(w, r)
		}
	})
	log.Println("Go API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
