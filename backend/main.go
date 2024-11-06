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

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")                   // Allow all origins
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS") // Allow specific HTTP methods
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow specific headers
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	mutex.Lock()
	if len(messages) == 0 {
		json.NewEncoder(w).Encode([]Message{}) // Return an empty array instead of null
	} else {
		json.NewEncoder(w).Encode(messages)
	}
	mutex.Unlock()
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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
