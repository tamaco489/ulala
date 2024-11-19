package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/miyabiii1210/ulala/go/model"
	"github.com/miyabiii1210/ulala/go/usecase"
)

type IAuthController interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	SignOut(c echo.Context) error
}

type authController struct {
	au usecase.IAuthUsecase
}

func NewAuthController(au usecase.IAuthUsecase) IAuthController {
	return &authController{au}
}

func (ac *authController) SignUp(c echo.Context) error {
	req := model.SignUpRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := ac.au.SignUp(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, res)
}

func (ac *authController) SignIn(c echo.Context) error {
	req := model.SignInRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := ac.au.SignIn(c, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (ac *authController) SignOut(c echo.Context) error {
	req := model.SignOutRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := ac.au.SignOut(c, req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
