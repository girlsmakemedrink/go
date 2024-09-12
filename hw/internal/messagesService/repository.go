package messagesService

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Message string `json: "Message"`
}

type MessageRepository interface {
	// CreateMessage - возвращаем созданное сообщение и ошибку
	CreateMessage(message Message) (Message, error)

	// GetAllMessages - Возвращаем массив из всех сообщений в БД и ошибку
	GetAllMessages() ([]Message, error)

	// UpdateMessageByID - Передаем id и Message, возвращаем обновленный Message и ошибку
	UpdateMessageByID(id int, message Message) (Message, error)

	// DeleteMessageByID - Передаем id для удаления, возвращаем удаленное сообщение и ошибку
	DeleteMessageByID(id int) (Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

// (r *messageRepository) привязывает данную функцию к нашему репозиторию

func (r *messageRepository) GetAllMessages() ([]Message, error) {
	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {
	createMessage := r.db.Create(&message)
	if createMessage.Error != nil {
		return Message{}, createMessage.Error
	}
	return message, nil
}

func (r *messageRepository) UpdateMessageByID(id int, message Message) (Message, error) {
	updatedMessage := r.db.First(&Message{}, id)
	if updatedMessage.Error != nil {
		return Message{}, updatedMessage.Error
	}

	updatedMessage = r.db.Model(&Message{}).Where("id = ?", id).Updates(&message)
	if updatedMessage.Error != nil {
		return Message{}, updatedMessage.Error
	}
	return message, nil
}

func (r *messageRepository) DeleteMessageByID(id int) (Message, error) {
	var message Message
	deletedMessage := r.db.First(&message, id)
	if deletedMessage.Error != nil {
		return Message{}, deletedMessage.Error
	}

	deletedMessage = r.db.Delete(&message, id)
	if deletedMessage.Error != nil {
		return  Message{}, deletedMessage.Error
	}
	return message, nil
}