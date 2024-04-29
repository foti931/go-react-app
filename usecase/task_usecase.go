package usecase

import (
	"go-rest-api/models"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]models.TaskResponse, error)
	GetTasksById(userId uint, taskId uint) (models.TaskResponse, error)
	CreateTask(task models.Task) (models.TaskResponse, error)
	UpdateTask(task models.Task, userId uint, taskId uint) (models.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) (models.TaskResponse, error)
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

// CreateTask implements ITaskUsecase.
func (t *taskUsecase) CreateTask(task models.Task) (models.TaskResponse, error) {
	if err := t.tv.TaskValidator(task); err != nil {
		return models.TaskResponse{}, err
	}
	if err := t.tr.CreateTask(&task); err != nil {
		return models.TaskResponse{}, err
	}

	response := models.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return response, nil
}

// DeleteTask implements ITaskUsecase.
func (t *taskUsecase) DeleteTask(userId uint, taskId uint) (models.TaskResponse, error) {
	if err := t.tr.DeleteTask(userId, taskId); err != nil {
		return models.TaskResponse{}, err
	}
	return models.TaskResponse{}, nil
}

// GetAllTasks implements ITaskUsecase.
func (t *taskUsecase) GetAllTasks(userId uint) ([]models.TaskResponse, error) {
	tasks := []models.Task{}
	if err := t.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []models.TaskResponse{}
	for _, v := range tasks {
		t := models.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}

	return resTasks, nil
}

// GetTasksById implements ITaskUsecase.
func (t *taskUsecase) GetTasksById(userId uint, taskId uint) (models.TaskResponse, error) {
	task := models.Task{}
	if err := t.tr.GetTasksById(&task, userId, taskId); err != nil {
		return models.TaskResponse{}, err
	}

	response := models.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return response, nil
}

// UpdateTask implements ITaskUsecase.
func (t *taskUsecase) UpdateTask(task models.Task, userId uint, taskId uint) (models.TaskResponse, error) {
	if err := t.tr.UpdateTask(&task, userId, taskId); err != nil {
		return models.TaskResponse{}, err
	}

	response := models.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return response, nil
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr: tr, tv: tv}
}
