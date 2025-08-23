package taskService

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     *Task
		mockSetup func(m *MockTaskRepository, input *Task)
		wantErr   bool
	}{
		{
			name:  "успешное создание задачи",
			input: &Task{Title: "Test", Completed: false},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("CreateTask", input).Return(input, nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании",
			input: &Task{Title: "Bad task", Completed: false},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("CreateTask", input).Return(&Task{}, errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.input)

			service := NewService(mockRepo)
			result, err := service.CreateTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetAllTask(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		want      []Task
		wantErr   bool
	}{
		{
			name: "успешное получение всех задач",
			mockSetup: func(m *MockTaskRepository) {
				tasks := []Task{
					{ID: 1, Title: "Task 1", Completed: false},
					{ID: 2, Title: "Task 2", Completed: true},
				}
				m.On("GetAllTask").Return(tasks, nil)
			},
			want: []Task{
				{ID: 1, Title: "Task 1", Completed: false},
				{ID: 2, Title: "Task 2", Completed: true},
			},
			wantErr: false,
		},
		{
			name: "ошибка при получении задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetAllTask").Return(nil, errors.New("db error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo)

			service := NewService(mockRepo)
			result, err := service.GetAllTask()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name      string
		id        uint
		input     *Task
		mockSetup func(m *MockTaskRepository, id uint, input *Task)
		want      *Task
		wantErr   bool
	}{
		{
			name:  "успешное обновление задачи",
			id:    1,
			input: &Task{ID: 1, Title: "Updated", Completed: true},
			mockSetup: func(m *MockTaskRepository, id uint, input *Task) {
				m.On("UpdateTaskByID", id, input).Return(input, nil)
			},
			want:    &Task{ID: 1, Title: "Updated", Completed: true},
			wantErr: false,
		},
		{
			name:  "ошибка при обновлении",
			id:    2,
			input: &Task{ID: 2, Title: "Fail", Completed: false},
			mockSetup: func(m *MockTaskRepository, id uint, input *Task) {
				m.On("UpdateTaskByID", id, mock.Anything).Return(nil, errors.New("update error"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id, tt.input)

			service := NewService(mockRepo)
			result, err := service.UpdateTaskByID(tt.id, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, result)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

func TestDeleteTaskByID(t *testing.T) {
	tests := []struct {
		name      string
		id        uint
		mockSetup func(m *MockTaskRepository, id uint)
		wantErr   bool
	}{
		{
			name: "успешно удалено",
			id:   1,
			mockSetup: func(m *MockTaskRepository, id uint) {
				m.On("DeleteTaskByID", id).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при удалении",
			id:   2,
			mockSetup: func(m *MockTaskRepository, id uint) {
				m.On("DeleteTaskByID", id).Return(errors.New("delete error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			tt.mockSetup(mockRepo, tt.id)

			service := NewService(mockRepo)
			err := service.DeleteTaskByID(tt.id)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}
