package handler

import (
	"app/api/application/interactor"
	"app/api/domain/service"
	"app/api/infrastructure/database"
	"app/api/infrastructure/repository"
)

type AppHandler struct {
	AuthHandler     AuthHandler
	UserHandler     UserHandler
	CategoryHandler CategoryHandler
}

func NewAppHandler(sqlHandler database.SQLHandler) *AppHandler {
	// repository
	userRepository := repository.NewUserRepository(sqlHandler)
	categoryRepository := repository.NewCategoryRepository(sqlHandler)

	// service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	categoryService := service.NewCategoryService(categoryRepository)

	// interactor
	userInteractor := interactor.NewUserInteractor(userService, authService)
	authInteractor := interactor.NewAuthInteractor(authService, userService)
	categoryInteractor := interactor.NewCategoryInteractor(categoryService)

	return &AppHandler{
		AuthHandler:     NewAuthHandler(authInteractor),
		UserHandler:     NewUserHandler(userInteractor),
		CategoryHandler: NewCategoryHandler(categoryInteractor),
	}
}
