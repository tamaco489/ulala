package validator

import (
	"errors"

	"github.com/miyabiii1210/ulala/go/model"
)

const Permisson string = "admin"

type IAdminValidator interface {
	AdminValidate(admin model.Admin) error
}

type adminValidator struct{}

func NewAdminValidator() IAdminValidator {
	return &adminValidator{}
}

func (adv *adminValidator) AdminValidate(admin model.Admin) error {
	if admin.Permisson != Permisson {
		return errors.New("You have insufficient privileges to access the service")
	}

	return nil
}
