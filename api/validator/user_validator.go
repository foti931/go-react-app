package validator

import (
	"go-rest-api/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IUserValidator interface {
	userValidator(user models.User) error
}

type userValidator struct{}

// userValidator implements IUserValidator.
func (u userValidator) userValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("limited max 30 char"),
			is.Email.Error("is not valid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("limited min 6 max 30 char"),
		),
	)
}

func NewUserValidator() IUserValidator {
	return userValidator{}
}
