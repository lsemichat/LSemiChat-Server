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
	TagHandler      TagHandler
	ThreadHandler   ThreadHandler
}

func NewAppHandler(sqlHandler database.SQLHandler) *AppHandler {
	// repository
	userRepository := repository.NewUserRepository(sqlHandler)
	categoryRepository := repository.NewCategoryRepository(sqlHandler)
	tagRepository := repository.NewTagRepository(sqlHandler)
	threadRepository := repository.NewThreadRepository(sqlHandler)

	// service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService()
	categoryService := service.NewCategoryService(categoryRepository)
	tagService := service.NewTagService(tagRepository)
	threadService := service.NewThreadService(threadRepository)

	// interactor
	userInteractor := interactor.NewUserInteractor(userService, authService)
	authInteractor := interactor.NewAuthInteractor(authService, userService)
	categoryInteractor := interactor.NewCategoryInteractor(categoryService)
	tagInteractor := interactor.NewTagInteractor(tagService, categoryService)
	threadInteractor := interactor.NewThreadInteractor(threadService, userService)

	return &AppHandler{
		AuthHandler:     NewAuthHandler(authInteractor),
		UserHandler:     NewUserHandler(userInteractor),
		CategoryHandler: NewCategoryHandler(categoryInteractor),
		TagHandler:      NewTagHandler(tagInteractor, categoryInteractor),
		ThreadHandler:   NewThreadHandler(threadInteractor),
	}
}
