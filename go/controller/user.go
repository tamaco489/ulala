package controller

import (
	"net/http"
	"strconv"

	"github.com/miyabiii1210/ulala/go/model"
	"github.com/miyabiii1210/ulala/go/usecase"

	"github.com/labstack/echo/v4"
)

// router側で使用するインターフェース
type IUserController interface {
	GetUsers(c echo.Context) error
	GetUserByUID(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

// mainで参照: usecaseで取得したレスポンスデータをechoを介してJSONフォーマットに変換するコンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) GetUsers(c echo.Context) error {
	res, err := uc.uu.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (uc *userController) GetUserByUID(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := uc.uu.GetUserByUID(uint32(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	req := model.UpdateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := uc.uu.UpdateUser(req, uint32(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := uc.uu.DeleteUser(uint32(uid))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusAccepted, res)
}
