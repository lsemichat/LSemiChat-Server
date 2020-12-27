package handler

import (
	"app/api/application/interactor"
	"app/api/infrastructure/lcontext"
	"app/api/infrastructure/lsession"
	"app/api/presentation/request"
	"app/api/presentation/response"
	"net/http"

	"github.com/pkg/errors"
)

type userHandler struct {
	userInteractor interactor.UserInteractor
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
	UpdateUserID(w http.ResponseWriter, r *http.Request)
	UpdatePassword(w http.ResponseWriter, r *http.Request)
	GetMe(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	DeleteMe(w http.ResponseWriter, r *http.Request)
	GetFollows(w http.ResponseWriter, r *http.Request)
	GetFollowers(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(ui interactor.UserInteractor) UserHandler {
	return &userHandler{
		userInteractor: ui,
	}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	src, err := ReadRequestBody(r, &request.UserCreateRequest{})
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to read request body"), "failed to read request")
		return
	}
	req, _ := src.(*request.UserCreateRequest)

	err = req.ValidateRequest(uh.userInteractor)
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to validation"), err.Error())
		return
	}

	user, err := uh.userInteractor.Create(req.UserID, req.Name, req.Mail, req.Image, req.Profile, req.Password)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to create user"), "failed to create user")
		return
	}

	response.Success(w, response.ConvertToUserResponse(user))
}

func (uh *userHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	userID, err := lcontext.GetUserIDFromContext(r.Context())
	if err != nil {
		response.Unauthorized(w, errors.Wrap(err, "failed to authentication"), "failed to authentication. please login")
		return
	}

	src, err := ReadRequestBody(r, &request.UserUpdateProfileRequest{})
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to read request body"), "failed to read request")
		return
	}
	req, _ := src.(*request.UserUpdateProfileRequest)
	if err = req.Validate(uh.userInteractor, userID); err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to validation"), err.Error())
		return
	}

	user, err := uh.userInteractor.UpdateProfile(userID, req.Name, req.Mail, req.Image, req.Profile)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to update"), "failed to update profile")
		return
	}
	response.Success(w, response.ConvertToUserResponse(user))
}

func (uh *userHandler) UpdateUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := lcontext.GetUserIDFromContext(r.Context())
	if err != nil {
		response.Unauthorized(w, errors.Wrap(err, "failed to authentication"), "failed to authentication. please login")
		return
	}

	src, err := ReadRequestBody(r, &request.UserUpdateUserIDRequest{})
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to read request body"), "failed to read request")
		return
	}
	req, _ := src.(*request.UserUpdateUserIDRequest)
	if err = req.Validate(uh.userInteractor, userID); err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to validation"), err.Error())
		return
	}

	user, err := uh.userInteractor.UpdateUserID(userID, req.UserID)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to update"), "failed to update userID")
		return
	}

	_, err = lsession.RestartSession(w, r, user.UserID)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to restart session"), "failed to restart session")
		return
	}
	response.Success(w, response.ConvertToUserResponse(user))
}

func (uh *userHandler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userID, err := lcontext.GetUserIDFromContext(r.Context())
	if err != nil {
		response.Unauthorized(w, errors.Wrap(err, "failed to authentication"), "failed to authentication. please login")
		return
	}

	src, err := ReadRequestBody(r, &request.UserUpdatePasswordRequest{})
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to read request body"), "failed to read request")
		return
	}
	req, _ := src.(*request.UserUpdatePasswordRequest)
	if err = req.Validate(); err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to validation"), err.Error())
		return
	}

	user, err := uh.userInteractor.UpdatePassword(userID, req.Password)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to update"), "failed to update userID")
		return
	}
	response.Success(w, response.ConvertToUserResponse(user))
}

func (uh *userHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	userID, err := lcontext.GetUserIDFromContext(r.Context())
	if err != nil {
		response.Unauthorized(w, errors.Wrap(err, "failed to authentication"), "failed to authentication")
		return
	}
	user, err := uh.userInteractor.GetByUserID(userID)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to get user"), "failed to get user")
		return
	}
	response.Success(w, response.ConvertToUserResponse(user))
}

func (uh *userHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := ReadPathParam(r, "id")
	if id == "" {
		response.BadRequest(w, errors.New("query param: id is empty"), "query param: id is empty")
		return
	}

	user, err := uh.userInteractor.GetByID(id)
	if err != nil {
		response.BadRequest(w, errors.Wrap(err, "failed to get user"), "failed to get user")
		return
	}
	response.Success(w, response.ConvertToUserResponse(user))
}

func (uh *userHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := uh.userInteractor.GetAll()
	if err != nil {
		response.BadRequest(w, err, "failed to get user")
		return
	}
	response.Success(w, users)
}

func (uh *userHandler) DeleteMe(w http.ResponseWriter, r *http.Request) {
	userID, err := lcontext.GetUserIDFromContext(r.Context())
	if err != nil {
		response.Unauthorized(w, errors.Wrap(err, "failed to authentication"), "failed to authentication")
		return
	}
	err = uh.userInteractor.Delete(userID)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to delete user"), "failed to delete")
		return
	}
	lsession.EndSession(w, r)
	response.NoContent(w)
}

func (uh *userHandler) GetFollows(w http.ResponseWriter, r *http.Request) {
	id := ReadPathParam(r, "id")
	users, err := uh.userInteractor.GetFollows(id)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to get follows"), "failed to get follows")
		return
	}
	response.Success(w, response.ConvertToUsersResponse(users))
}

func (uh *userHandler) GetFollowers(w http.ResponseWriter, r *http.Request) {
	id := ReadPathParam(r, "id")
	users, err := uh.userInteractor.GetFollowers(id)
	if err != nil {
		response.InternalServerError(w, errors.Wrap(err, "failed to get followers"), "failed to get followers")
		return
	}
	response.Success(w, response.ConvertToUsersResponse(users))
}
