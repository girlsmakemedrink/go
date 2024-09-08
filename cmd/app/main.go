package main

import (
	"github.com/gorilla/mux"
	"hw/internal/database"
	"hw/internal/handlers"
	"hw/internal/messagesService"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/delete", handler.DeleteMessageHandler).Methods("DELETE")
	router.HandleFunc("/api/patch", handler.PatchMessageHandler).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}