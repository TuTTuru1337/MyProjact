package taskService

type TaskRepository interface {
	CreateTask(task *Task) (*Task, error)
	GetAllTask() ([]Task, error)
	UpdateTaskByID(id uint, task *Task) (*Task, error)
	DeleteTaskByID(id uint) error
}

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *Task) (*Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTask()
}

func (s *TaskService) UpdateTaskByID(id uint, task *Task) (*Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
