package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	//"log"
	//"fmt"
)

var message string
var dbRecord Message

//Сообщение в POST запросе
type requestBody struct {
	Message string `json:"message"`
	ID int `json:"ID"`
}

//POST Выводим сообщение из тела запроса
func SendMessage(w http.ResponseWriter, r *http.Request) {
	var rb requestBody
	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Создаем запись в БД с заполнением поля text
	sendMessage := Message{Text: rb.Message}
	result := DB.Create(&sendMessage)
	if result.Error != nil {
		http.Error(w, "Failed to create message", http.StatusBadRequest)
		return
	}
	// Устанавливаем заголовок и возвращаем джейсон
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sendMessage)
}

//GET Выводим все записи из таблицы
func ShowMessages(w http.ResponseWriter, r *http.Request) {
	var records []Message
	showMessage := DB.Find(&records)

	if showMessage.Error != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

//DELETE Удаляем запись по ID
func DeleteMessage (w http.ResponseWriter, r *http.Request) {
	var rb requestBody
	err := json.NewDecoder(r.Body).Decode(&rb)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := DB.First(&dbRecord, rb.ID)
	if result.Error != nil {
		http.Error(w, "Message Not found", http.StatusNotFound)
		return
	}

	result = DB.Delete(&dbRecord, &rb)
	if result.Error != nil {
		http.Error(w, "Error deleting message", http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dbRecord)

}

//PATCH Обновляем запись по ID
	func UpdateMessage (w http.ResponseWriter, r *http.Request) {
		var rb requestBody
		err := json.NewDecoder(r.Body).Decode(&rb)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		DB.Model(&dbRecord).Where("id = ?", rb.ID).Updates(Message{Text: rb.Message})

		result := DB.First(&dbRecord, rb.ID)
		if result.Error != nil {
			http.Error(w, "Message Not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dbRecord)

	}

	func main() {
		InitDB()

		DB.AutoMigrate(&dbRecord)

		router := mux.NewRouter()
		router.HandleFunc("/api/show", ShowMessages).Methods("GET")
		router.HandleFunc("/api/send", SendMessage).Methods("POST")
		router.HandleFunc("/api/delete", DeleteMessage).Methods("DELETE")
		router.HandleFunc("/api/update", UpdateMessage).Methods("PATCH")
		http.ListenAndServe(":8080", router)
	}