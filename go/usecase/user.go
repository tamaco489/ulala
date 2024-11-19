package usecase

import (
	"os"

	"github.com/miyabiii1210/ulala/go/model"
	"github.com/miyabiii1210/ulala/go/repository"
	"github.com/miyabiii1210/ulala/go/validator"
)

// controller側で使用するインターフェース
type IUserUsecase interface {
	GetUsers() ([]model.GetUserResponse, error)
	GetUserByUID(uid uint32) (model.GetUserResponse, error)
	UpdateUser(req model.UpdateUserRequest, uid uint32) (model.UpdateUserResponse, error)
	DeleteUser(uid uint32) (model.DefaultResponse, error)
}

type userUsecase struct {
	ur  repository.IUserRepository
	uv  validator.IUserValidator
	adv validator.IAdminValidator
}

// mainで参照: repositoryで取得したデータオブジェクトを指定したレスポンスフォーマットに変換するコンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator, adv validator.IAdminValidator) IUserUsecase {
	return &userUsecase{ur, uv, adv}
}

func (uu *userUsecase) GetUsers() ([]model.GetUserResponse, error) {
	admin := model.Admin{Permisson: os.Getenv("PERMISSION")}
	if err := uu.adv.AdminValidate(admin); err != nil {
		return nil, err
	}

	users := []model.User{}
	if err := uu.ur.GetUsers(&users); err != nil {
		return []model.GetUserResponse{}, err
	}

	res := []model.GetUserResponse{}
	for _, u := range users {
		u := model.GetUserResponse{
			UID:   u.UID,
			Name:  u.Name,
			Email: u.Email,
		}
		res = append(res, u)
	}

	return res, nil
}

func (uu *userUsecase) GetUserByUID(uid uint32) (model.GetUserResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserByUID(&user, uid); err != nil {
		return model.GetUserResponse{}, err
	}

	res := model.GetUserResponse{
		UID:   user.UID,
		Name:  user.Name,
		Email: user.Email,
	}

	return res, nil
}

func (uu *userUsecase) UpdateUser(req model.UpdateUserRequest, uid uint32) (model.UpdateUserResponse, error) {
	if err := uu.uv.UpdateUserValidate(req); err != nil {
		return model.UpdateUserResponse{}, err
	}

	u := model.User{Name: req.Name, Email: req.Email}
	if err := uu.ur.UpdateUser(&u, uid); err != nil {
		return model.UpdateUserResponse{}, err
	}

	res := model.UpdateUserResponse{Name: u.Name, Email: u.Email}

	return res, nil
}

func (uu *userUsecase) DeleteUser(uid uint32) (model.DefaultResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserByUID(&user, uid); err != nil {
		return model.DefaultResponse{}, err
	}

	if err := uu.ur.DeleteUser(uid); err != nil {
		return model.DefaultResponse{}, err
	}

	return model.DefaultResponse{
		Message: model.DEFAULT_RESPONSE_202_MESSAGE,
	}, nil
}
