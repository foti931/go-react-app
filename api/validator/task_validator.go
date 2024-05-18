package validator

import (
	"go-rest-api/models"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ITaskValidator interface {
	TaskValidator(Task models.Task) error
}

type TaskValidator struct{}

// TaskValidator implements ITaskValidator.
func (u TaskValidator) TaskValidator(Task models.Task) error {
	return validation.ValidateStruct(&Task,
		validation.Field(
			&Task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
		),
	)
}

func NewTaskValidator() ITaskValidator {
	return TaskValidator{}
}
