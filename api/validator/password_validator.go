package validator

import (
	"go-rest-api/models"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type IPasswordResetValidator interface {
	passwordResetValidator(password models.PasswordReset) error
}

type passwordResetValidator struct{}

func (p passwordResetValidator) passwordResetValidator(password models.PasswordReset) error {
	return validation.ValidateStruct(&password,
		validation.Field(
			&password.Email,
			validation.Required.Error("email is required"),
			is.Email.Error("invalid email format"),
		),
	)
}

func NewPasswordResetValidator() IPasswordResetValidator {
	return &passwordResetValidator{}
}
