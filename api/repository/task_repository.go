package repository

import (
	"fmt"
	"go-rest-api/models"

	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]models.Task, userId uint) error
	GetTasksById(task *models.Task, userId uint, taskId uint) error
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task, userId uint, taskId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

// CreateTask implements ITaskRepository.
func (t *TaskRepository) CreateTask(task *models.Task) error {
	if err := t.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

// DeleteTask implements ITaskRepository.
func (t *TaskRepository) DeleteTask(userId uint, taskId uint) error {
	result := t.db.Where("id=? AND user_id=?", taskId, userId).Delete(&models.Task{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("task id: %d not exist", taskId)
	}

	return nil
}

// GetAllTasks implements ITaskRepository.
func (t *TaskRepository) GetAllTasks(tasks *[]models.Task, userId uint) error {
	if err := t.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

// GetTasksById implements ITaskRepository.
func (t *TaskRepository) GetTasksById(task *models.Task, userId uint, taskId uint) error {
	if err := t.db.Joins("User").Where("tasks.id=? AND tasks.user_id=?", taskId, userId).First(task).Error; err != nil {
		return err
	}

	return nil
}

// UpdateTask implements ITaskRepository.
func (t *TaskRepository) UpdateTask(task *models.Task, userId uint, taskId uint) error {
	result := t.db.Model(&task).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("task id: %d not exist", taskId)
	}

	return nil
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db: db}
}
