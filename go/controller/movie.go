package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/miyabiii1210/ulala/go/usecase"
)

type IMovieController interface {
	GetMovieCategories(c echo.Context) error
	GetMovieListByTypeID(c echo.Context) error
	GetMovie(c echo.Context) error
}

type movieController struct {
	mu usecase.IMovieUsecase
}

func NewMovieController(mu usecase.IMovieUsecase) IMovieController {
	return &movieController{mu}
}

func (mc *movieController) GetMovieCategories(c echo.Context) error {
	res, err := mc.mu.GetMovieCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (mc *movieController) GetMovieListByTypeID(c echo.Context) error {
	typeID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := mc.mu.GetMovieListByTypeID(uint32(typeID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (mc *movieController) GetMovie(c echo.Context) error {
	movieID, err := strconv.Atoi(c.Param("movie_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := mc.mu.GetMovie(uint32(movieID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
