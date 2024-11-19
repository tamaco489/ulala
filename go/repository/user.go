package repository

import (
	"github.com/miyabiii1210/ulala/go/model"

	"gorm.io/gorm"
)

// usecase側で使用するインターフェース
type IUserRepository interface {
	CreateUser(user *model.User) error
	GetUsers(users *[]model.User) error
	GetUserByUID(user *model.User, uid uint32) error
	UpdateUser(user *model.User, uid uint32) error
	DeleteUser(uid uint32) error
}

type userRepository struct {
	db *gorm.DB
}

// mainで参照: DBへの問い合わせ結果をgorm構造体に反映して返却するコンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return IUserRepository(&userRepository{db})
}

func (ur *userRepository) CreateUser(user *model.User) error {
	return ur.db.Create(&user).Error
}

func (ur *userRepository) GetUsers(users *[]model.User) error {
	return ur.db.Find(&users).Error
}

func (ur *userRepository) GetUserByUID(user *model.User, uid uint32) error {
	return ur.db.First(&user, uid).Error
}

func (ur *userRepository) UpdateUser(user *model.User, uid uint32) error {
	return ur.db.Where("uid = ?", uid).Updates(model.User{Name: user.Name, Email: user.Email}).Error
}

func (ur *userRepository) DeleteUser(uid uint32) error {
	return ur.db.Where("uid = ?", uid).Delete(&model.User{}).Error
}
