package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var message string

//Structure
type requestBody struct {
	Message string `json:"message"`
}

//POST
func BodyHandler(w http.ResponseWriter, r *http.Request) {
	var rb requestBody

	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		http.Error(w, "Invalid request body!!", http.StatusBadRequest)
		return
	}

	message = rb.Message

	// Устанавливаем заголовок и возвращаем джейсон
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rb)
}
//GET
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %v", message)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/body", BodyHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
