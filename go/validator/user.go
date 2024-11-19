package validator

import (
	"github.com/miyabiii1210/ulala/go/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

const (
	INVALID_NAME_LENGTH  string = "User name must be between 1 and 30 characters in length"
	INVALID_EMAIL_LENGTH string = "Email address must be between 1 and 50 characters in length"
	INVALID_EMAIL_FORMAT string = "Email address is invalid format"
	NOT_EXIST_EMAIL      string = "Email address is required"
)

type IUserValidator interface {
	CreateUserValidate(req model.User) error
	UpdateUserValidate(req model.UpdateUserRequest) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) CreateUserValidate(req model.User) error {
	return validation.ValidateStruct(&req,
		validation.Field(
			&req.Email,
			validation.Required.Error(NOT_EXIST_EMAIL),
			validation.RuneLength(1, 50).Error(INVALID_EMAIL_LENGTH),
			is.Email.Error(INVALID_EMAIL_FORMAT),
		),
	)
}

func (uv *userValidator) UpdateUserValidate(req model.UpdateUserRequest) error {
	switch {
	case req.Name != "":
		if err := validation.ValidateStruct(&req,
			validation.Field(
				&req.Name,
				validation.Length(1, 30).Error(INVALID_NAME_LENGTH),
			),
		); err != nil {
			return err
		}

	case req.Email != "":
		if err := validation.ValidateStruct(&req,
			validation.Field(
				&req.Email,
				validation.RuneLength(1, 50).Error(INVALID_EMAIL_LENGTH),
				is.Email.Error(INVALID_EMAIL_FORMAT),
			),
		); err != nil {
			return err
		}

	default:
		return nil
	}

	return nil
}
