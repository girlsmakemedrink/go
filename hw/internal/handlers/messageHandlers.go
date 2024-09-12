package handlers

import (
	"encoding/json"
	"hw/internal/messagesService" // Импортируем наш сервис
	"net/http"
)

type Handler struct {
	Service *messagesService.MessageService
}

// Нужна для создания структуры Handler на этапе инициализации приложения
func NewHandler(service *messagesService.MessageService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) ErrorHandler(w http.ResponseWriter, statusCode int, message string) {
	http.Error(w, message, statusCode)
}

func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	messages, err := h.Service.GetAllMessages()
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при получении сообщений")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при получении сообщения")
		return
	}

	createdMessage, err := h.Service.CreateMessage(message)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при создании сообщения")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdMessage)
}

func (h *Handler) PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при получении сообщения")
		return
	}

	updatedMessage, err := h.Service.UpdateMessageByID(int(message.ID),message)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при обновлении сообщения")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedMessage)

}

func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	var message messagesService.Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при получении сообщения")
		return
	}

	deletedMessage, err := h.Service.DeleteMessageByID(int(message.ID))
	if err != nil {
		h.ErrorHandler(w, http.StatusInternalServerError, "Ошибка при удалении сообщения")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedMessage)
}
