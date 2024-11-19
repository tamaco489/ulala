package repository

import (
	"github.com/miyabiii1210/ulala/go/model"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	SignUp(auth *model.UserFirebaseAuthentication) error
	SignIn(firebaseUID string, user *model.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
	return IAuthRepository(&authRepository{db})
}

func (ar *authRepository) SignUp(auth *model.UserFirebaseAuthentication) error {
	return ar.db.Create(auth).Error
}

func (ar *authRepository) SignIn(firebaseUID string, user *model.User) error {
	// SELECT u.uid FROM users AS u JOIN user_firebase_authentications AS fa ON u.auth_uuid = fa.uuid WHERE fa.firebase_uid = 'jRO1p7ot7meV7bm20Ol1HULmDGO2';
	if err := ar.db.Table("users").
		Select("users.uid").
		Joins("JOIN user_firebase_authentications AS fa ON users.auth_uuid = fa.uuid").
		Where("fa.firebase_uid = ?", firebaseUID).
		First(&user).
		Error; err != nil {
		return err
	}

	return nil
}
