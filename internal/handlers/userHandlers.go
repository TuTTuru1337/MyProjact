package handlers

import (
	"Tutturu/internal/models"
	"Tutturu/internal/userService/service"
	"Tutturu/internal/web/users"
	"context"
)

type UserHandler struct {
	service *service.UserServiceImpl
}

func NewUserHandler(svc *service.UserServiceImpl) *UserHandler {
	return &UserHandler{
		service: svc,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.service.GetAllUsers(ctx)
	if err != nil {
		return users.GetUsers500JSONResponse{
			Message: "Internal server error",
		}, nil
	}

	response := users.GetUsers200JSONResponse{}
	for _, u := range allUsers {
		user := users.User{
			Id:        &u.ID,
			Email:     (*users.Email)(&u.Email),
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
		}
		if u.DeletedAt != nil {
			user.DeletedAt = u.DeletedAt
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body

	email := string(userRequest.Email)

	userToCreate := models.User{
		Email:    email,
		Password: userRequest.Password,
	}

	createdUser, err := h.service.CreateUser(ctx, userToCreate)
	if err != nil {
		return users.PostUsers500JSONResponse{
			Message: "Failed to create user",
		}, nil
	}

	emailResponse := users.Email(createdUser.Email)

	response := users.PostUsers201JSONResponse{
		Id:        &createdUser.ID,
		Email:     &emailResponse,
		CreatedAt: &createdUser.CreatedAt,
		UpdatedAt: &createdUser.UpdatedAt,
	}
	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, req users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := req.Id
	body := req.Body

	user, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		return users.PatchUsersId404Response{}, nil
	}

	if body.Email != "" {
		user.Email = string(body.Email)
	}
	if body.Password != "" {
		user.Password = body.Password
	}

	updatedUser, err := h.service.UpdateUser(ctx, *user)
	if err != nil {
		return users.PatchUsersId500JSONResponse{
			Message: "Failed to update user",
		}, nil
	}

	emailResponse := users.Email(updatedUser.Email)

	return users.PatchUsersId200JSONResponse{
		Id:        &updatedUser.ID,
		Email:     &emailResponse,
		CreatedAt: &updatedUser.CreatedAt,
		UpdatedAt: &updatedUser.UpdatedAt,
	}, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, req users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	err := h.service.DeleteUser(ctx, req.Id)
	if err != nil {
		return users.DeleteUsersId404Response{}, nil
	}
	return users.DeleteUsersId204Response{}, nil
}
