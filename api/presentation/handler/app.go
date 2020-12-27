package handler

import (
	"app/api/application/interactor"
	"app/api/domain/service"
	"app/api/infrastructure/database"
	"app/api/infrastructure/repository"
)

type AppHandler struct {
	AuthHandler AuthHandler
	UserHandler UserHandler
}

func NewAppHandler(sqlHandler database.SQLHandler) *AppHandler {
	// repository
	userRepository := repository.NewUserRepository(sqlHandler)

	// service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()

	// interactor
	userInteractor := interactor.NewUserInteractor(userService, authService)
	authInteractor := interactor.NewAuthInteractor(authService, userService)

	return &AppHandler{
		AuthHandler: NewAuthHandler(authInteractor),
		UserHandler: NewUserHandler(userInteractor),
	}
}
