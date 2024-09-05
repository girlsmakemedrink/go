package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

var message string

//Сообщение в POST запросе
type requestBody struct {
	Message string `json:"message"`
}

//POST Выводим сообщение из тела запроса
func BodyHandler(w http.ResponseWriter, r *http.Request) {
	var rb requestBody
	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		http.Error(w, "Error BodyHandler", http.StatusBadRequest)
		return
	}

	// Создаем запись в БД с заполнением поля text
	record := Message{Text: rb.Message}
	result := DB.Create(&record)
	if result.Error != nil {
		log.Fatal("failed to create record", result.Error)
	}
	// Устанавливаем заголовок и возвращаем джейсон
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rb)
}

//GET Выводим все записи из таблицы
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var records []Message
	record := DB.Find(&records)
	if record.Error != nil {
		http.Error(w, "Error HelloHandler", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	router.HandleFunc("/api/body", BodyHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}