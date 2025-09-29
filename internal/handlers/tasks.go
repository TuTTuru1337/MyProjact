package handlers

import (
	"Tutturu/internal/models"
	"Tutturu/internal/service"
	"Tutturu/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *TaskHandler {
	return &TaskHandler{
		service: svc,
	}
}

// GetTasks implements tasks.ServerInterface.
func (h *TaskHandler) GetTasks(c echo.Context) error {
	allTasks, err := h.service.GetAllTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]tasks.Task, len(allTasks))
	for i, t := range allTasks {
		response[i] = tasks.Task{
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

// PostTasks implements tasks.ServerInterface.
func (h *TaskHandler) PostTasks(c echo.Context) error {
	var taskReq tasks.TaskRequest
	if err := c.Bind(&taskReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	taskToCreate := models.Task{
		Task:   taskReq.Task,
		IsDone: *taskReq.IsDone,
		UserID: taskReq.UserId,
	}

	createdTask, err := h.service.CreateTask(c.Request().Context(), taskToCreate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := tasks.Task{
		Id:        &createdTask.ID,
		UserId:    &createdTask.UserID,
		Task:      &createdTask.Task,
		IsDone:    &createdTask.IsDone,
		CreatedAt: &createdTask.CreatedAt,
		UpdatedAt: &createdTask.UpdatedAt,
	}

	return c.JSON(http.StatusCreated, response)
}

// GetTasksByUserId implements tasks.ServerInterface.
func (h *TaskHandler) GetTasksByUserId(c echo.Context, userId uint) error {
	tasksList, err := h.service.GetTasksByUserID(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := make([]tasks.Task, len(tasksList))
	for i, t := range tasksList {
		response[i] = tasks.Task{
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

// DeleteTasksId implements tasks.ServerInterface.
func (h *TaskHandler) DeleteTasksId(c echo.Context, id uint) error {
	err := h.service.DeleteTask(c.Request().Context(), int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TaskHandler) PatchTasksId(c echo.Context, id uint) error {
	var taskReq tasks.TaskRequest
	if err := c.Bind(&taskReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	task, err := h.service.GetTaskByID(c.Request().Context(), int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Task not found"})
	}

	// Обновляем только переданные поля
	task.Task = taskReq.Task
	if taskReq.IsDone != nil {
		task.IsDone = *taskReq.IsDone
	}
	task.UserID = taskReq.UserId

	updatedTask, err := h.service.UpdateTask(c.Request().Context(), *task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	response := tasks.Task{
		Id:        &updatedTask.ID,
		UserId:    &updatedTask.UserID,
		Task:      &updatedTask.Task,
		IsDone:    &updatedTask.IsDone,
		CreatedAt: &updatedTask.CreatedAt,
		UpdatedAt: &updatedTask.UpdatedAt,
	}

	return c.JSON(http.StatusOK, response)
}
