package main

import (
	"fmt"

	"github.com/miyabiii1210/ulala/go/controller"
	"github.com/miyabiii1210/ulala/go/datastore"
	"github.com/miyabiii1210/ulala/go/repository"
	"github.com/miyabiii1210/ulala/go/router"
	"github.com/miyabiii1210/ulala/go/usecase"
	"github.com/miyabiii1210/ulala/go/validator"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func handler(db *gorm.DB) *echo.Echo {
	adminValidator := validator.NewAdminValidator()

	// user
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator, adminValidator)
	userController := controller.NewUserController(userUsecase)

	// auth
	authRepository := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository, userValidator)
	authController := controller.NewAuthController(authUsecase)

	// movie
	movieRepositorty := repository.NewMovieRepository(db)
	movieUsecase := usecase.NewMovieUsecase(movieRepositorty)
	movieController := controller.NewMovieController(movieUsecase)

	e := router.NewRouter(userController, authController, movieController)
	return e
}

func main() {
	fmt.Println("[debug] starting api server...")
	db := datastore.NewDBConnection()
	e := handler(db)
	if e == nil {
		fmt.Println("failed to start api server")
		return
	}
	defer db.Rollback()
	e.Logger.Fatal(e.Start(":8080"))
}
