package handlers

import (
	"context"
	"errors"
	"hw/internal/usersService" // Импортируем наш сервис
	"hw/internal/web/users"    // Импортируем пакет users
)
type UserHandler struct {
	Service *usersService.UserService
}

// Нужна для создания структуры UserHandler на этапе инициализации приложения
func NewUserHandler (service *usersService.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	// Получение всех пользователей из сервиса
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	// Создаем переменную респон типа 200джейсонРеспонс
	// Которую мы потом передадим в качестве ответа
	response := users.GetUsers200JSONResponse{}

	// Заполняем слайс response всеми пользователями из БД
	for _, usr := range allUsers {
		user := users.User{
			Id:   &usr.ID,
			User: &usr.User,
		}
		response = append(response, user)
	}

	// САМОЕ ПРЕКРАСНОЕ. Возвращаем просто респонс и nil!
	return response, nil
}

func (h *UserHandler) PostUsers(_ context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и создаем пользователя
	userToCreate := usersService.User{User: *userRequest.User}
	createdUser, err := h.Service.CreateUser(userToCreate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PostUsers201JSONResponse{
		Id:   &createdUser.ID,
		User: &createdUser.User,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) PatchUsers(_ context.Context, request users.PatchUsersRequestObject) (users.PatchUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	// Обращаемся к сервису и обновляем пользователя
	userToUpdate := usersService.User{User: *userRequest.User}
	updatedUser, err := h.Service.UpdateUserByID(*userRequest.Id,userToUpdate)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.PatchUsers200JSONResponse{
		Id:   &updatedUser.ID,
		User: &updatedUser.User,
	}
	// Просто возвращаем респонс!
	return response, nil
}

func (h *UserHandler) DeleteUsers(_ context.Context, request users.DeleteUsersRequestObject) (users.DeleteUsersResponseObject, error) {
	// Распаковываем тело запроса напрямую, без декодера!
	userRequest := request.Body
	if userRequest == nil {
		return nil, errors.New("userRequest is nil")
	}

	// Обращаемся к сервису и удаляем пользователя
	deletedUser, err := h.Service.DeleteUserByID(*userRequest.Id)

	if err != nil {
		return nil, err
	}
	// создаем структуру респонс
	response := users.DeleteUsers200JSONResponse{
		Id:   &deletedUser.ID,
		User: &deletedUser.User,
	}
	// Просто возвращаем респонс!
	return response, nil
}