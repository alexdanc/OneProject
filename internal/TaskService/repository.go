package TaskService

import "gorm.io/gorm"

type RequestBodyRepository interface {
	CreateTask(task RequestBody) error
	GetAllTasks() ([]RequestBody, error)
	GetTaskByID(id string) (RequestBody, error)
	UpdateTask(task RequestBody) error
	DeleteTaskByID(id string) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) RequestBodyRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task RequestBody) error {

	return r.db.Create(&task).Error
}

func (r *TaskRepository) GetAllTasks() ([]RequestBody, error) {
	var tasks []RequestBody
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id string) (RequestBody, error) {
	var tas RequestBody
	err := r.db.First(&tas, id).Error
	return tas, err

}

func (r *TaskRepository) UpdateTask(task RequestBody) error {
	return r.db.Save(&task).Error
}

func (r *TaskRepository) DeleteTaskByID(id string) error {
	return r.db.Delete(&RequestBody{}, "id = ?", id).Error
}
