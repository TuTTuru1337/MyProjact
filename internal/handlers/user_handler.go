package handlers

import (
	"Tutturu/internal/models"
	service2 "Tutturu/internal/service"
	"Tutturu/internal/userService/service"
	"Tutturu/internal/web/users"
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type UserHandler struct {
	service     *service.UserServiceImpl
	taskService *service2.Service
}

func NewUserHandler(svc *service.UserServiceImpl, taskService *service2.Service) *UserHandler {
	return &UserHandler{
		service:     svc,
		taskService: taskService,
	}
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	allUsers, err := h.service.GetAllUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]users.User, len(allUsers))
	for i, u := range allUsers {
		email := openapi_types.Email(u.Email)
		response[i] = users.User{
			Id:        &u.ID,
			Email:     &email,
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
		}
		if u.DeletedAt != nil {
			response[i].DeletedAt = u.DeletedAt
		}
	}
	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) GetUsersId(c echo.Context, id uint) error {
	user, err := h.service.GetUserByID(c.Request().Context(), int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	email := openapi_types.Email(user.Email) // Альтернативное решение

	return c.JSON(http.StatusOK, users.User{
		Id:        &user.ID,
		Email:     &email,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	})
}

func (h *UserHandler) GetUsersIdTasks(c echo.Context, id uint) error {
	tasksList, err := h.taskService.GetTasksByUserID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]users.Task, len(tasksList))
	for i, t := range tasksList {
		response[i] = users.Task{
			Id:        &t.ID,
			UserId:    &t.UserID,
			Task:      &t.Task,
			IsDone:    &t.IsDone,
			CreatedAt: &t.CreatedAt,
			UpdatedAt: &t.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, response)
}

func (h *UserHandler) PostUsers(c echo.Context) error {
	var userReq users.UserRequest
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	userToCreate := models.User{
		Email:    string(userReq.Email),
		Password: userReq.Password,
	}

	createdUser, err := h.service.CreateUser(c.Request().Context(), userToCreate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	email := openapi_types.Email(createdUser.Email) // Альтернативное решение

	return c.JSON(http.StatusCreated, users.User{
		Id:        &createdUser.ID,
		Email:     &email,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
	})
}

func (h *UserHandler) DeleteUsersId(c echo.Context, id uint) error {
	err := h.service.DeleteUser(c.Request().Context(), int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHandler) PatchUsersId(c echo.Context, id uint) error {
	var userReq users.UserRequest
	if err := c.Bind(&userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, err := h.service.GetUserByID(c.Request().Context(), int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	if userReq.Email != "" {
		user.Email = string(userReq.Email)
	}
	if userReq.Password != "" {
		user.Password = userReq.Password
	}

	updatedUser, err := h.service.UpdateUser(c.Request().Context(), *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	email := openapi_types.Email(updatedUser.Email)

	return c.JSON(http.StatusOK, users.User{
		Id:        &updatedUser.ID,
		Email:     &email,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
	})
}
