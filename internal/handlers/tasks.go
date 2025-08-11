package handlers

import (
	"Tutturu/internal/models"
	"Tutturu/internal/service"
	"Tutturu/internal/web/tasks"
	"context"
)

type TaskHandler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *TaskHandler {
	return &TaskHandler{
		service: svc,
	}
}

func (h *TaskHandler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks(context.Background())
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, t := range allTasks {
		task := tasks.Task{
			Id:     &t.ID,
			Task:   &t.Task,
			IsDone: &t.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := models.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.service.CreateTask(ctx, taskToCreate)
	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, req tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := req.Id
	body := req.Body

	task, err := h.service.GetTaskByID(ctx, id)
	if err != nil {
		return tasks.PatchTasksId404Response{}, nil
	}

	if body.Task != nil {
		task.Task = *body.Task
	}
	if body.IsDone != nil {
		task.IsDone = *body.IsDone
	}

	updatedTask, err := h.service.UpdateTask(ctx, *task)
	if err != nil {
		return nil, err
	}

	return tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}, nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	err := h.service.DeleteTask(ctx, req.Id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, nil
	}

	return tasks.DeleteTasksId204Response{}, nil
}
