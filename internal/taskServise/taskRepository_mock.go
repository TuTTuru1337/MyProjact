package taskService

import (
	"github.com/stretchr/testify/mock"
)

type Task struct {
	ID        uint
	Title     string
	Completed bool
}

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) CreateTask(task *Task) (*Task, error) {
	args := m.Called(task)
	var t *Task
	if res := args.Get(0); res != nil {
		t = res.(*Task)
	}
	return t, args.Error(1)
}

func (m *MockTaskRepository) GetAllTask() ([]Task, error) {
	args := m.Called()
	var tasks []Task
	if res := args.Get(0); res != nil {
		tasks = res.([]Task)
	}
	return tasks, args.Error(1)
}

func (m *MockTaskRepository) UpdateTaskByID(id uint, task *Task) (*Task, error) {
	args := m.Called(id, task)
	var t *Task
	if res := args.Get(0); res != nil {
		t = res.(*Task)
	}
	return t, args.Error(1)
}

func (m *MockTaskRepository) DeleteTaskByID(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
